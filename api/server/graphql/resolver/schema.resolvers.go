package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	api "github.com/portey/twitch-streaming/api/gen/service"
	"github.com/portey/twitch-streaming/api/server/graphql/generated"
	"github.com/portey/twitch-streaming/api/server/graphql/model"
	log "github.com/sirupsen/logrus"
)

func (r *mutationResolver) SubscribeToStreamer(ctx context.Context, name string) (bool, error) {
	_, err := r.client.SubscribeToStreamer(ctx, &api.SubscribeToStreamerRequest{
		StreamerName: name,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) GetSignedURL(ctx context.Context) (string, error) {
	res, err := r.client.GetSignUrl(ctx, &api.GetSignUrlRequest{})
	if err != nil {
		return "", err
	}

	return res.Url, nil
}

func (r *queryResolver) EventsAggregatedByStreamer(ctx context.Context) ([]*model.EventByStreamer, error) {
	res, err := r.client.GetEventCountsAggregatedByStreamer(ctx, &api.GetEventCountsAggregatedByStreamerRequest{})
	if err != nil {
		return nil, err
	}

	list := make([]*model.EventByStreamer, len(res.Items))
	for i, item := range res.Items {
		list[i] = &model.EventByStreamer{
			Streamer: item.StreamerName,
			Count:    int(item.Count),
		}
	}

	return list, nil
}

func (r *queryResolver) ExchangeToken(ctx context.Context, accessToken string) (string, error) {
	res, err := r.client.ExchangeToken(ctx, &api.ExchangeTokenRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		return "", err
	}

	return res.AccessToken, nil
}

func (r *subscriptionResolver) Events(ctx context.Context, streamerName string) (<-chan *model.Event, error) {
	ch := make(chan *model.Event, 0)
	id, err := r.mediator.Subscribe(streamerName, func(event *model.Event) error {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
			return nil
		case ch <- event:
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		if err := r.mediator.Unsubscribe(streamerName, id); err != nil {
			log.WithError(err).Error("web socket unsubscribe error")
		}
	}()

	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
