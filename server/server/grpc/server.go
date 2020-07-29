package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	api "github.com/portey/twitch-streaming/server/gen/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func New(port int, service Service) (*Server, error) {
	// build Service
	srv := Server{
		addr: fmt.Sprintf(":%d", port),
		server: grpc.NewServer(
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(log.NewEntry(log.StandardLogger())),
				grpc_recovery.UnaryServerInterceptor(),
				func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
					token := metautils.ExtractIncoming(ctx).Get("Authorization")

					return handler(context.WithValue(ctx, "Authorization", token), req) // nolint
				},
			)),
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_logrus.StreamServerInterceptor(log.NewEntry(log.StandardLogger())),
				grpc_recovery.StreamServerInterceptor(),
			)),
		),
	}
	api.RegisterTwitchStreamServiceServer(srv.server, newResolver(service))

	return &srv, nil
}

type Server struct {
	addr      string
	server    *grpc.Server
	runErr    error
	readiness bool
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	log.Info("grpc srv: begin run")

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		s.runErr = err
		log.Error("grpc srv: run error: ", err)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := s.server.Serve(lis)
		log.Error("grpc srv: end run > ", err)
		s.runErr = err
	}()

	go func() {
		<-ctx.Done()
		s.server.GracefulStop()
		log.Info("grpc srv: graceful stop")
	}()

	s.readiness = true
}

func (s *Server) HealthCheck() error {
	if !s.readiness {
		return errors.New("grpc service is't ready yet")
	}
	if s.runErr != nil {
		return errors.New("grpc service: run issue")
	}
	return nil
}
