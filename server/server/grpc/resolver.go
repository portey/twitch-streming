package grpc

import (
	"context"

	api "github.com/portey/twitch-streaming/server/gen/service"
	"github.com/portey/twitch-streaming/server/models"
)

type Service interface {
	GetSignUrl(ctx context.Context) (string, error)
	ExchangeToken(ctx context.Context, token string) (string, error)
	SubscribeToStreamer(ctx context.Context, name string) error
	GetEventsCountAggregatedByStreamer(ctx context.Context) ([]models.EventsCountsByStreamer, error)
	GetEventsCountAggregatedByStreamerAndType(ctx context.Context) ([]models.EventsCountsByStreamerAndType, error)
}

type Resolver struct {
	service Service
}

func newResolver(service Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

func (r *Resolver) GetSignUrl(ctx context.Context, _ *api.GetSignUrlRequest) (*api.GetSignUrlResponse, error) {
	url, err := r.service.GetSignUrl(ctx)
	if err != nil {
		return nil, err
	}

	return &api.GetSignUrlResponse{Url: url}, nil
}

func (r *Resolver) ExchangeToken(ctx context.Context, request *api.ExchangeTokenRequest) (*api.ExchangeTokenResponse, error) {
	res, err := r.service.ExchangeToken(ctx, request.GetAccessToken())
	if err != nil {
		return nil, err
	}

	return &api.ExchangeTokenResponse{
		AccessToken: res,
	}, nil
}

func (r *Resolver) SubscribeToStreamer(ctx context.Context, request *api.SubscribeToStreamerRequest) (*api.SubscribeToStreamerResponse, error) {
	err := r.service.SubscribeToStreamer(ctx, request.GetStreamerName())
	return &api.SubscribeToStreamerResponse{}, err
}

func (r *Resolver) GetEventCountsAggregatedByStreamer(ctx context.Context, _ *api.GetEventCountsAggregatedByStreamerRequest) (*api.GetEventCountsAggregatedByStreamerResponse, error) {
	list, err := r.service.GetEventsCountAggregatedByStreamer(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*api.GetEventCountsAggregatedByStreamerResponse_Item, len(list))
	for i := range list {
		item := list[i]
		res[i] = &api.GetEventCountsAggregatedByStreamerResponse_Item{
			StreamerName: item.StreamerName,
			Count:        int32(item.Count),
		}
	}

	return &api.GetEventCountsAggregatedByStreamerResponse{
		Items: res,
	}, nil
}

func (r *Resolver) GetEventCountsAggregatedByStreamerAndType(ctx context.Context, _ *api.GetEventCountsAggregatedByStreamerAndTypeRequest) (*api.GetEventCountsAggregatedByStreamerAndTypeResponse, error) {
	list, err := r.service.GetEventsCountAggregatedByStreamerAndType(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*api.GetEventCountsAggregatedByStreamerAndTypeResponse_Item, len(list))
	for i := range list {
		item := list[i]
		res[i] = &api.GetEventCountsAggregatedByStreamerAndTypeResponse_Item{
			StreamerName: item.StreamerName,
			EventType:    string(item.EventType),
			Count:        int32(item.Count),
		}
	}

	return &api.GetEventCountsAggregatedByStreamerAndTypeResponse{
		Items: res,
	}, nil
}
