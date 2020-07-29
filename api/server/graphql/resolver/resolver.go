package resolver

import (
	api "github.com/portey/twitch-streaming/api/gen/service"
	"github.com/portey/twitch-streaming/api/server/graphql/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type EvenMediator interface {
	Subscribe(streamerName string, handler func(event *model.Event) error) (string, error)
	Unsubscribe(streamerName, id string) error
}

type Resolver struct {
	client   api.TwitchStreamServiceClient
	mediator EvenMediator
}

func New(cfg Config, mediator EvenMediator) (*Resolver, error) {
	conn, err := createConnection(cfg)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		client:   api.NewTwitchStreamServiceClient(conn),
		mediator: mediator,
	}, nil
}
