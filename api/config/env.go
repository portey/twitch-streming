package config

import (
	"github.com/portey/twitch-streaming/api/redis"
	"github.com/portey/twitch-streaming/api/server/graphql/resolver"
	"github.com/spf13/viper"
)

func Read() Config {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")

	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("GRAPH_QL_PORT", "8080")
	viper.SetDefault("HEALTH_CHECK_PORT", "8889")

	viper.SetDefault("REDIS_ADDRESS", "localhost:6379")
	viper.SetDefault("REDIS_POOL_SIZE", "10")
	viper.SetDefault("REDIS_PUB_SUB_CHANNEL", "events")

	viper.SetDefault("SERVER_HOST", "localhost")
	viper.SetDefault("SERVER_PORT", "8090")
	viper.SetDefault("SERVER_TIMEOUT", "10s")

	return Config{
		LogLevel:          viper.GetString("LOG_LEVEL"),
		GraphQLServerPort: viper.GetInt("GRAPH_QL_PORT"),
		HealthCheckPort:   viper.GetInt("HEALTH_CHECK_PORT"),
		RedisCfg: redis.Config{
			Addr:     viper.GetString("REDIS_ADDRESS"),
			PoolSize: viper.GetInt("REDIS_POOL_SIZE"),
			Channel:  viper.GetString("REDIS_PUB_SUB_CHANNEL"),
		},
		ServerConnCfg: resolver.Config{
			Host:    viper.GetString("SERVER_HOST"),
			Port:    viper.GetInt("SERVER_PORT"),
			Timeout: viper.GetDuration("SERVER_TIMEOUT"),
		},
	}
}
