package config

import (
	"github.com/portey/twitch-streaming/server/redis"
	"github.com/portey/twitch-streaming/server/service/twitch"
	"github.com/portey/twitch-streaming/server/storage/postgres"
	"github.com/spf13/viper"
)

func Read() Config {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")

	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("GRPC_PORT", "8090")
	viper.SetDefault("HTTP_PORT", "8081")
	viper.SetDefault("HEALTH_CHECK_PORT", "8888")

	viper.SetDefault("JWT_LIVE_TIME", "86400s")

	viper.SetDefault("REDIS_ADDRESS", "localhost:6379")
	viper.SetDefault("REDIS_POOL_SIZE", "10")
	viper.SetDefault("REDIS_PUB_SUB_CHANNEL", "events")

	viper.SetDefault("POSTGRES_ADDRESS", "postgres://secret:secret@localhost:5432/events?sslmode=disable")

	viper.SetDefault("TWITCH_EVENT_SECRET", "secret-key")

	return Config{
		LogLevel:        viper.GetString("LOG_LEVEL"),
		GRPCServerPort:  viper.GetInt("GRPC_PORT"),
		HTTPServerPort:  viper.GetInt("HTTP_PORT"),
		HealthCheckPort: viper.GetInt("HEALTH_CHECK_PORT"),

		JWTPrivateKey: []byte(viper.GetString("JWT_PRIVATE_KEY")),
		JWTLiveTime:   viper.GetDuration("JWT_LIVE_TIME"),

		RedisCfg: redis.Config{
			Addr:     viper.GetString("REDIS_ADDRESS"),
			PoolSize: viper.GetInt("REDIS_POOL_SIZE"),
			Channel:  viper.GetString("REDIS_PUB_SUB_CHANNEL"),
		},
		PostgresCfg: postgres.Config{
			Address: viper.GetString("POSTGRES_ADDRESS"),
		},
		TwitchCfg: twitch.Config{
			ClientID:     viper.GetString("TWITCH_AUTH_CLIENT_ID"),
			ClientSecret: viper.GetString("TWITCH_AUTH_CLIENT_SECRET"),
			RedirectURL:  viper.GetString("TWITCH_AUTH_REDIRECT_URL"),
			CallbackURL:  viper.GetString("TWITCH_EVENT_CALLBACK_URL"),
			Secret:       viper.GetString("TWITCH_EVENT_SECRET"),
		},
	}
}
