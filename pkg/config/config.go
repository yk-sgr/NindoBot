package config

import (
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	IsDev      bool   `split_words:"true" default:"false"`
	SentryDSN  string `split_words:"true"`
	BotToken   string `envconfig:"TOKEN" required:"true"`
	BotPrefix  string `envconfig:"PREFIX" default:"nd!"`
	ShardCount uint   `split_words:"true" default:"1"`
	ShardIDs   []uint `envconfig:"SHARD_IDS" default:"0"`
}

func Load() *Config {
	loadEnvFile()

	var conf Config
	err := envconfig.Process("BOT", &conf)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("Could not load config.", zap.Error(err))
	}
	return &conf
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		zap.L().Warn("Could not load .env file.", zap.Error(err))
	}
}
