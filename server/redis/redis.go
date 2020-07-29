package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type (
	Config struct {
		Addr     string
		PoolSize int
		Channel  string
	}

	Sink struct {
		client  *redis.Client
		channel string
	}

	Handler func(ctx context.Context, payload []byte) error
)

var errClientNotInitialized = errors.New("redis client: not initialized yet")

// New returns the initialized Sink object.
func NewRedis(ctx context.Context, cfg Config) (*Sink, error) {
	log.Infof("PubSub init: addr=%s", cfg.Addr)
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		PoolSize: cfg.PoolSize,
	})

	go func() {
		<-ctx.Done()
		err := client.Close()
		if err != nil {
			log.WithError(err).Error("close redis connection")
			return
		}
		log.Info("close redis connection")
	}()

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Sink{
		client:  client,
		channel: cfg.Channel,
	}, nil
}

func (o *Sink) Sink(ctx context.Context, payload []byte) error {
	return o.client.Publish(ctx, o.channel, payload).Err()
}

func (o *Sink) HealthCheck() error {
	if o.client == nil {
		return errClientNotInitialized
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return o.client.Ping(ctx).Err()
}
