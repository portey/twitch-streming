package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func New(port int, service Service) (*Server, error) {
	// build http server
	httpSrv := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	// build Server
	var srv Server
	srv.setupHTTP(&httpSrv)
	srv.service = service

	return &srv, nil
}

type Server struct {
	http      *http.Server
	runErr    error
	readiness bool
	service   Service
}

func (s *Server) setupHTTP(srv *http.Server) {
	srv.Handler = s.buildHandler()
	s.http = srv
}

func (s *Server) buildHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/event", s.trackEvent).Methods(http.MethodPost)

	return r
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	log.Info("http service: begin run")

	go func() {
		defer wg.Done()
		log.Debug("http service: addr=", s.http.Addr)
		err := s.http.ListenAndServe()
		s.runErr = err
		log.Info("http service: end run > ", err)
	}()

	go func() {
		<-ctx.Done()
		sdCtx, _ := context.WithTimeout(context.Background(), 5*time.Second) // nolint
		err := s.http.Shutdown(sdCtx)
		if err != nil {
			log.Info("http service shutdown (", err, ")")
		}
	}()

	s.readiness = true
}

func (s *Server) HealthCheck() error {
	if !s.readiness {
		return errors.New("http service is't ready yet")
	}
	if s.runErr != nil {
		return errors.New("http service: run issue")
	}
	return nil
}
