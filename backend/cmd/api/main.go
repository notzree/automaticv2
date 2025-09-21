package main

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/notzree/automaticv2/v2/config"
	"go.uber.org/fx"
)

type Server struct {
	Echo         *echo.Echo
	ServerConfig *config.HTTPServerConfig
}

func NewServer(lc fx.Lifecycle, config *config.Config) *Server {
	server := &Server{
		Echo:         echo.New(),
		ServerConfig: &config.HTTPServer,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.StartServer()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.StopServer(ctx)
		},
	})

	return server
}

func (s *Server) StartServer() error {
	slog.Info("Starting server", "address", s.ServerConfig.HTTPAddress)
	return s.Echo.Start(s.ServerConfig.HTTPAddress)
}

func (s *Server) StopServer(ctx context.Context) error {
	slog.Info("Stopping server")
	return s.Echo.Shutdown(ctx)
}

func (s *Server) RegisterRoutes() {
	s.Echo.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// TODO: Add more routes
}

var HTTPModule = fx.Options(
	fx.Provide(
		NewServer,
		config.NewConfig,
	),
	fx.Invoke(func(s *Server) {
		s.RegisterRoutes()
	}),
)

func main() {
	fx.New(HTTPModule).Run()
}
