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

	Observer struct {
		client  *redis.Client
		ps      *redis.PubSub
		channel string
	}

	Handler func(ctx context.Context, payload []byte) error
)

var errClientNotInitialized = errors.New("redis client: not initialized yet")

// New returns the initialized Observer object.
func NewRedis(ctx context.Context, cfg Config) (*Observer, error) {
	log.Infof("Cache init: addr=%s", cfg.Addr)
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

	ps := client.Subscribe(ctx, cfg.Channel)
	if _, err := ps.Receive(ctx); err != nil {
		return nil, err
	}

	return &Observer{
		client:  client,
		ps:      ps,
		channel: cfg.Channel,
	}, nil
}

func (o *Observer) Sub(ctx context.Context, handler Handler) {
	go func() {
		for e := range o.ps.Channel() {
			log.Debug("redis: received a message")
			if err := handler(ctx, []byte(e.Payload)); err != nil {
				log.WithError(err).Error("handling pub/sub")
			}
		}
		log.Debug("redis: exited")
	}()
}

func (o *Observer) HealthCheck() error {
	if o.ps == nil || o.client == nil {
		return errClientNotInitialized
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := o.ps.Ping(ctx); err != nil {
		return err
	}

	return o.client.Ping(ctx).Err()
}
