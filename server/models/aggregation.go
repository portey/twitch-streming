package models

type EventsCountsByStreamer struct {
	StreamerName string `db:"streamer_name"`
	Count        int    `db:"count"`
}

type EventsCountsByStreamerAndType struct {
	EventType EventType `db:"event_type"`
	EventsCountsByStreamer
}
