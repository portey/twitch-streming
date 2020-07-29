package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/portey/twitch-streaming/server/config"
	"github.com/portey/twitch-streaming/server/jwt"
	"github.com/portey/twitch-streaming/server/redis"
	"github.com/portey/twitch-streaming/server/server/grpc"
	"github.com/portey/twitch-streaming/server/server/healthcheck"
	"github.com/portey/twitch-streaming/server/server/http"
	"github.com/portey/twitch-streaming/server/service"
	"github.com/portey/twitch-streaming/server/service/twitch"
	"github.com/portey/twitch-streaming/server/storage/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	// read service cfg from os env
	cfg := config.Read()

	// init logger
	initLogger(cfg.LogLevel)

	log.Info("Service starting...")

	// prepare main context
	ctx, cancel := context.WithCancel(context.Background())
	setupGracefulShutdown(cancel)

	keyData, _ := pem.Decode(cfg.JWTPrivateKey)
	if keyData == nil {
		log.Fatal("invalid private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(keyData.Bytes)
	if err != nil {
		log.WithError(err).Fatal("parse private key")
	}

	jwpPtovider := jwt.NewProvider(privateKey, cfg.JWTLiveTime)

	// initializing redis
	sink, err := redis.NewRedis(ctx, cfg.RedisCfg)
	if err != nil {
		log.WithError(err).Fatal("redis initialization")
	}

	eventsRepo, err := postgres.New(cfg.PostgresCfg)
	if err != nil {
		log.WithError(err).Fatal("postgres initialization")
	}

	logicService := service.New(eventsRepo, sink, jwpPtovider, twitch.New(cfg.TwitchCfg))

	grpcServer, err := grpc.New(cfg.GRPCServerPort, logicService)
	if err != nil {
		log.WithError(err).Fatal("graphql server init")
	}

	httpServer, err := http.New(cfg.HTTPServerPort, logicService)
	if err != nil {
		log.WithError(err).Fatal("http server init")
	}

	healthChecks := []func() error{
		grpcServer.HealthCheck,
		httpServer.HealthCheck,
		sink.HealthCheck,
	}

	// build health check srv
	healthSrv := healthcheck.New(cfg.HealthCheckPort, healthChecks...)

	var wg = &sync.WaitGroup{}

	// run srv
	grpcServer.Run(ctx, wg)
	httpServer.Run(ctx, wg)
	healthSrv.Run(ctx, wg)

	// wait while services work
	wg.Wait()
	log.Info("Service stopped")
}

func initLogger(logLevel string) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	switch strings.ToLower(logLevel) {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}

func setupGracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Error("Got Interrupt signal")
		stop()
	}()
}
