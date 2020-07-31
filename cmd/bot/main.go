package main

import (
	"github.com/ForYaSee/NindoBot/pkg/bot"
	"github.com/ForYaSee/NindoBot/pkg/config"
	"github.com/ForYaSee/NindoBot/pkg/logger"
	"github.com/ForYaSee/NindoBot/pkg/web"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
)

func main() {
	conf := config.Load()
	logger.Init(conf.IsDev)
	if err := sentry.Init(sentry.ClientOptions{Dsn: conf.SentryDSN}); err != nil {
		zap.L().Error("Could not initialize sentry.", zap.Error(err))
	}

	b := bot.New(bot.Config{
		Token:      conf.BotToken,
		ShardCount: conf.ShardCount,
		ShardIDs:   conf.ShardIDs,
		Prefix:     conf.BotPrefix,
	})
	go func() {
		zap.L().Info("Bot started.")
		zap.L().Fatal("Error while serving bot.", zap.Error(b.Start()))
	}()

	w := web.New()
	w.Run(conf.WebAddress)
}
