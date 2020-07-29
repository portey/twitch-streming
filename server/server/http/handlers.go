package http

import (
	"context"
	"io"
	"net/http"
)

type Service interface {
	TrackEvent(ctx context.Context, data io.Reader) error
}

func (s *Server) trackEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if err := s.service.TrackEvent(r.Context(), r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
