package models

import "time"

const (
	EventTypeFollow    EventType = "FOLLOW"
	EventTypeSubscribe EventType = "SUBSCRIBE"
)

// Event type
type EventType string

// Model to track single event
type Event struct {
	ID             string    `json:"id"`
	ActionUserName string    `json:"action_user_name"`
	StreamerName   string    `json:"streamer_name"`
	EventType      EventType `json:"event_type"`
	CreatedAt      time.Time `json:"created_at"`
}
