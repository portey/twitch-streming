package mediator

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/portey/twitch-streaming/api/server/graphql/model"
	log "github.com/sirupsen/logrus"
)

type Mediator struct {
	listeners map[string]map[string]struct{}
	handlers  map[string]func(event *model.Event) error
	mx        sync.RWMutex
}

func New() *Mediator {
	return &Mediator{
		listeners: make(map[string]map[string]struct{}),
		handlers:  make(map[string]func(event *model.Event) error),
	}
}

func (m *Mediator) Subscribe(streamerName string, f func(event *model.Event) error) (string, error) {
	id := uuid.New().String()

	m.mx.Lock()
	defer m.mx.Unlock()

	if _, ok := m.listeners[streamerName]; !ok {
		m.listeners[streamerName] = make(map[string]struct{}, 1)
	}

	m.listeners[streamerName][id] = struct{}{}
	m.handlers[id] = f

	return id, nil
}

func (m *Mediator) Unsubscribe(streamerName, id string) error {
	m.mx.Lock()
	defer m.mx.Unlock()

	if _, ok := m.listeners[streamerName]; ok {
		delete(m.listeners[streamerName], id)
	}
	delete(m.handlers, id)

	return nil
}

func (m *Mediator) Handle(_ context.Context, payload []byte) error {
	var event struct {
		ID             string    `json:"id"`
		ActionUserName string    `json:"action_user_name"`
		StreamerName   string    `json:"streamer_name"`
		EventType      string    `json:"event_type"`
		CreatedAt      time.Time `json:"created_at"`
	}

	if err := json.Unmarshal(payload, &event); err != nil {
		log.WithError(err).Error("invalid mediator error")
	}

	m.mx.RLock()
	defer m.mx.RUnlock()
	if handlers, ok := m.listeners[event.StreamerName]; ok {
		for id := range handlers {
			if handler, ok2 := m.handlers[id]; ok2 {
				eventType := model.EventType(event.EventType)

				err := handler(&model.Event{
					ID:             event.ID,
					SubscriberName: event.ActionUserName,
					SubscribedTo:   event.StreamerName,
					Type:           &eventType,
				})
				if err != nil {
					log.WithError(err).Error("sending event to handler")
				}
			}
		}
	}

	return nil
}
