package config

import (
	"time"

	"github.com/portey/twitch-streaming/server/redis"
	"github.com/portey/twitch-streaming/server/service/twitch"
	"github.com/portey/twitch-streaming/server/storage/postgres"
)

type Config struct {
	LogLevel        string
	HealthCheckPort int
	GRPCServerPort  int
	HTTPServerPort  int

	JWTPrivateKey []byte
	JWTLiveTime   time.Duration

	RedisCfg    redis.Config
	PostgresCfg postgres.Config
	TwitchCfg   twitch.Config
}
