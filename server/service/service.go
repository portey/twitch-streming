package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/portey/twitch-streaming/server/models"
	"github.com/portey/twitch-streaming/server/service/twitch"
	log "github.com/sirupsen/logrus"
)

type Storage interface {
	GetEventsCountAggregatedByStreamer(ctx context.Context) ([]models.EventsCountsByStreamer, error)
	GetEventsCountAggregatedByStreamerAndType(ctx context.Context) ([]models.EventsCountsByStreamerAndType, error)
	SaveEvent(ctx context.Context, event models.Event) error
}

type RealTimeEventSink interface {
	Sink(ctx context.Context, payload []byte) error
}

type AuthProvider interface {
	GenerateToken(ctx context.Context, viewerName string) (string, error)
	IsAuthorized(ctx context.Context) bool
}

type Service struct {
	storage      Storage
	sink         RealTimeEventSink
	authProvider AuthProvider
	client       *twitch.Client
}

func New(
	storage Storage,
	sink RealTimeEventSink,
	authProvider AuthProvider,
	client *twitch.Client,
) *Service {
	return &Service{
		storage:      storage,
		sink:         sink,
		authProvider: authProvider,
		client:       client,
	}
}

// Get sign url to redirect user to login view
func (s *Service) GetSignUrl(_ context.Context) (string, error) {
	return s.client.GetSignedUrl(), nil
}

// Exchange auth2 code to project access token
func (s *Service) ExchangeToken(ctx context.Context, clientAccessToken string) (string, error) {
	userID, err := s.client.GetCurrentUser(ctx, clientAccessToken)
	if err != nil {
		return "", err
	}

	return s.authProvider.GenerateToken(ctx, userID)
}

// Subscribe to streamer
func (s *Service) SubscribeToStreamer(ctx context.Context, name string) error {
	if !s.authProvider.IsAuthorized(ctx) {
		return models.ErrNotAuthorized
	}

	return s.client.SubscribeToUserFollows(ctx, name)
}

// Get events count aggregated by user
func (s *Service) GetEventsCountAggregatedByStreamer(ctx context.Context) ([]models.EventsCountsByStreamer, error) {
	if !s.authProvider.IsAuthorized(ctx) {
		return nil, models.ErrNotAuthorized
	}

	return s.storage.GetEventsCountAggregatedByStreamer(ctx)
}

// Get events count aggregated by user and type
func (s *Service) GetEventsCountAggregatedByStreamerAndType(ctx context.Context) ([]models.EventsCountsByStreamerAndType, error) {
	if !s.authProvider.IsAuthorized(ctx) {
		return nil, models.ErrNotAuthorized
	}
	return s.storage.GetEventsCountAggregatedByStreamerAndType(ctx)

}

// Track new event from twitch, callback
func (s *Service) TrackEvent(ctx context.Context, data io.Reader) error {
	var req struct {
		Data []struct {
			FromName string `json:"from_id"`
			ToName   string `json:"to_name"`
		} `json:"data"`
	}

	if err := json.NewDecoder(data).Decode(&req); err != nil {
		return err
	}

	if len(req.Data) != 1 {
		return errors.New("not valid data length")
	}

	event := models.Event{
		ID:             uuid.New().String(),
		ActionUserName: req.Data[0].FromName,
		StreamerName:   req.Data[0].ToName,
		EventType:      models.EventTypeFollow,
		CreatedAt:      time.Now(),
	}

	if err := s.storage.SaveEvent(ctx, event); err != nil {
		return err
	}

	eventData, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	if err := s.sink.Sink(ctx, eventData); err != nil {
		log.WithError(err).Error("sink error")
	}

	return nil
}
