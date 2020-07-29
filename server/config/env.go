package config

import (
	"github.com/portey/twitch-streaming/server/redis"
	"github.com/portey/twitch-streaming/server/service/twitch"
	"github.com/portey/twitch-streaming/server/storage/postgres"
	"github.com/spf13/viper"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA1ZDOfs1Am0nXfHHU2qloFn4Ak7FsdvsADv3Mw+rMraIkLvy+
ewWQEKQztKPz82bVITotrg3M6zBTUtIGusQEHynAdL/LZTyqS8u3Ji3EV2/zEUXL
JKl1XH0KSkaXlunKDtSPgziqy/UZDfVFXSU+9elVO9KmLwr2b9YK+RHIOA/YvkzJ
eSr17GNJHa5194tw8OiGgInIm2HUmQ06ffJueMZuqcK3n5JTOC0UJTlu7xhJAg64
9vEJ0o9TCQzGmriJvXvYKu/9RUvPFyAmscLzqGPuKBQU/No63sW5DNbr766KvbN9
2oR9ZmuMlInIj80d+jE03YO5qBzu3m/Mgss/dQIDAQABAoIBABWa7qj5Tr3m/HvE
cEomtTcBTEGkw1pODjV2C48OY+J08sAmJYcAixCD7A8sPvOyxYD1B8YB82cfnm5U
fQXL6rbUXHnzJTp6pqfAVijwnFpUIvzuWJy+3/aZV78n4RINWLmjW1llFDJJ+1zk
IT5JFFOAZXUF3HYJ3+B0gOutltnE4UUw5mgjhQPKuegZgpxxqUAx/3OZghj0XpSi
BkgtLjfURZN3GLqSJk+2C3RGihGDok3jGsij2ZvWfh6XuTncDLg7mx1U2qUOLRcO
NaHRu5/nUoTbCj7dZHYiow6x7sFsm0AS6kSmDgc4GihY8OCx+xWCCE223kErigHH
HjbFgkECgYEA9VUWSe6QM6v8wgptO3ZPy4vDZhla7AbHwdH6jNIBBA0CbEAf5J5i
m5pVlC51J+dccSUXT0zfDIL2qwkA11BjBv48wrjSme/09XT3khepR0YwMEoONZSb
+Yjt/z7yI8awITKwN6jdvzyq/2XNZTPVD9UDM7FpURsnmN8v+PLjZo0CgYEA3tob
3wc/gwf1jzqpdFkQzFJojrIN0jbpA9WlI25+LmDbqcTP/kNpSvd8aCqVspop0XjU
LCBG4PtDUxaaMI+1tdviXLAgoeOONvNZSGFJkv2ecy8gPbPIb7A6/aIlcLB4eoGV
wZ2tIFgBNnEbU8YerRfuYPDTmDJhtsk5zZ6+VokCgYEAseYAYp4WfRGzGHX9HYg+
dqgjig7KyqpkAd0k5SaHTAuu/RZGyj473P++HNTPaZ8wfm9aBswEVgtFmWLO9FM7
a9/B6aWiObQyGie2On7j5mY0HWAmC552uC0d/+ACMAUsxPX/qGzQV5NDoC2PElrS
nkdlqf91EjYxsX3uSITAdyUCgYEA2vf+yM9rZYmNjDW8yVi2e71BhWyIzhQsMxta
zwzDqTh8vjxnjtOYAxRYRlYJj1uRWYTbHZx9aJUa2upriOm8RzwOGLrq5Ycddvr3
sHn/fBH9/fnBOT+M48mKvSr0lNyhFOZ9Sqhus0glsOPEUTVrcPMBxHj9wB9JCfyA
8nxXU9kCgYEA4ZRiZ7iFI4nQIHyau0oOIOLf27UljV5AHVAUs2+W4pSG+ETpC1W5
/vODz+sLWFPDhP9s+bU56qGlKslM6Etyqi/5uEg8gm17mNozaGhF75m4yaL6oPhv
ObL0fLzoBYfeRFkYeIYxeCY8y4J2YzOBwX9xNNfQCRBwjfHhvy6rrSU=
-----END RSA PRIVATE KEY-----`

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

	//todo remove this envs
	viper.SetDefault("TWITCH_AUTH_CLIENT_ID", "4zhcatnjekdh2hzo49iwcgi6dlxyd9")
	viper.SetDefault("TWITCH_AUTH_CLIENT_SECRET", "vsjfqe7iwdoc3zf1adnjco6konexrx")
	viper.SetDefault("TWITCH_AUTH_REDIRECT_URL", "http://localhost")
	viper.SetDefault("TWITCH_EVENT_CALLBACK_URL", "http://localhost/track")
	viper.SetDefault("JWT_PRIVATE_KEY", privateKey)

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
