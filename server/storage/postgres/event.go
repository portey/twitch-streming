package postgres

import (
	"context"

	"github.com/portey/twitch-streaming/server/models"
)

func (r *EventsRepo) GetEventsCountAggregatedByStreamer(ctx context.Context) ([]models.EventsCountsByStreamer, error) {
	const query = `
SELECT streamer_name,count(*)
FROM events
GROUP BY streamer_name
`

	var dest []models.EventsCountsByStreamer
	if err := r.db.SelectContext(ctx, &dest, query); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *EventsRepo) GetEventsCountAggregatedByStreamerAndType(ctx context.Context) ([]models.EventsCountsByStreamerAndType, error) {
	const query = `
SELECT streamer_name, event_type, count(*)
FROM events
GROUP BY streamer_name, event_type
`

	var dest []models.EventsCountsByStreamerAndType
	if err := r.db.SelectContext(ctx, &dest, query); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *EventsRepo) SaveEvent(ctx context.Context, event models.Event) error {
	const query = `
INSERT INTO 
	events(id, streamer_name, event_type, created_at) 
VALUES 
	(:id, :streamer_name, :event_type, :created_at);
`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":            event.ID,
		"streamer_name": event.StreamerName,
		"event_type":    event.EventType,
		"created_at":    event.CreatedAt,
	})
	return err
}
