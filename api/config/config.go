package config

import (
	"github.com/portey/twitch-streaming/api/redis"
	"github.com/portey/twitch-streaming/api/server/graphql/resolver"
)

type Config struct {
	LogLevel          string
	HealthCheckPort   int
	GraphQLServerPort int
	RedisCfg          redis.Config
	ServerConnCfg     resolver.Config
}
