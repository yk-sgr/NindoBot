package web

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/adaptor"
	"github.com/gofiber/fiber"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type WebServer struct {
	app *fiber.App
}

func New() *WebServer {
	app := fiber.New()

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	return &WebServer{
		app: app,
	}
}

func (s *WebServer) Run(address string) {
	err := s.app.Listen(address)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("Error while serving Webserver.", zap.Error(err))
	}
}
