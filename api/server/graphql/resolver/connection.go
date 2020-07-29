package resolver

import (
	"context"
	"fmt"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func createConnection(config Config) (*grpc.ClientConn, error) {
	timeout, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	return grpc.DialContext(
		timeout,
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_logrus.UnaryClientInterceptor(log.NewEntry(log.StandardLogger())),
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_logrus.StreamClientInterceptor(log.NewEntry(log.StandardLogger())),
		)),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}
