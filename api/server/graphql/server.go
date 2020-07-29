package graphql

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "github.com/99designs/gqlgen/cmd" // nolint
	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/portey/twitch-streaming/api/server/graphql/generated"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func New(port int, resolver generated.ResolverRoot) (*Server, error) {
	// build http server
	httpSrv := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	// build Server
	var srv Server
	srv.resolver = resolver
	srv.setupHTTP(&httpSrv)

	return &srv, nil
}

type Server struct {
	http      *http.Server
	runErr    error
	readiness bool
	resolver  generated.ResolverRoot
}

func (s *Server) setupHTTP(srv *http.Server) {
	srv.Handler = s.buildHandler()
	s.http = srv
}

func (s *Server) buildHandler() http.Handler {
	srv := gqlhandler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: s.resolver,
	}))

	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	handler := http.NewServeMux()

	handler.Handle("/query", authMiddleware(srv))
	handler.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	return cors.AllowAll().Handler(handler)
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	log.Info("gql service: begin run")

	go func() {
		defer wg.Done()
		log.Debug("gql service: addr=", s.http.Addr)
		err := s.http.ListenAndServe()
		s.runErr = err
		log.Info("gql service: end run > ", err)
	}()

	go func() {
		<-ctx.Done()
		sdCtx, _ := context.WithTimeout(context.Background(), 5*time.Second) // nolint
		err := s.http.Shutdown(sdCtx)
		if err != nil {
			log.Info("gql service shutdown (", err, ")")
		}
	}()

	s.readiness = true
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if token := req.Header.Get("Authorization"); token != "" {
			ctx := req.Context()
			// note: possible conflict with context value key
			ctx = context.WithValue(ctx, "Authorization", token) // nolint
			ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", token)

			req = req.WithContext(ctx)

		}

		next.ServeHTTP(w, req)
	})
}

func (s *Server) HealthCheck() error {
	if !s.readiness {
		return errors.New("gql service is't ready yet")
	}
	if s.runErr != nil {
		return errors.New("gql service: run issue")
	}
	return nil
}
