package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/portey/twitch-streaming/api/config"
	"github.com/portey/twitch-streaming/api/mediator"
	"github.com/portey/twitch-streaming/api/redis"
	"github.com/portey/twitch-streaming/api/server/graphql"
	"github.com/portey/twitch-streaming/api/server/graphql/resolver"
	"github.com/portey/twitch-streaming/api/server/healthcheck"
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

	pubsub, err := redis.NewRedis(ctx, cfg.RedisCfg)
	if err != nil {
		log.WithError(err).Fatal("redis initialization")
	}

	med := mediator.New()
	pubsub.Sub(ctx, med.Handle)

	graphQLResolver, err := resolver.New(cfg.ServerConnCfg, med)
	if err != nil {
		log.WithError(err).Fatal("graphql server resolver creation")
	}

	graphQlServer, err := graphql.New(cfg.GraphQLServerPort, graphQLResolver)
	if err != nil {
		log.WithError(err).Fatal("graphql server init")
	}

	healthChecks := []func() error{
		graphQlServer.HealthCheck,
		pubsub.HealthCheck,
	}

	// build health check srv
	healthSrv := healthcheck.New(cfg.HealthCheckPort, healthChecks...)

	var wg = &sync.WaitGroup{}

	// run srv
	graphQlServer.Run(ctx, wg)
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
