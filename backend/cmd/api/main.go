package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/notzree/automaticv2/v2/config"
	"github.com/notzree/automaticv2/v2/pkg/errors"
	automatic_http "github.com/notzree/automaticv2/v2/pkg/http"
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
	RegisterRoutes(s.Echo)
}

// RegisterRoutes sets up all routes using Echo's built-in grouping
func RegisterRoutes(e *echo.Echo) {
	// Global middleware
	e.Use(automatic_http.LoggingMiddleware)
	e.Use(automatic_http.CORSMiddleware)

	// Health check (no group)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "healthy"})
	})

	// API v1 group
	api := e.Group("/api/v1")

	// Public routes
	public := api.Group("/public")
	automatic_http.RegisterHandlers(public, ListPublicRoutes())

	// Protected routes
	protected := api.Group("/protected")
	protected.Use(automatic_http.RequireAuth)
	automatic_http.RegisterHandlers(protected, ListUserRoutes())

	// Admin routes (nested under protected)
	admin := protected.Group("/admin")
	admin.Use(automatic_http.RequireAdmin)
	automatic_http.RegisterHandlers(admin, ListAdminRoutes())

	// Agent routes (example of domain-specific grouping)
	agents := api.Group("/agents")
	agents.Use(automatic_http.RequireAuth)
	automatic_http.RegisterHandlers(agents, ListAgentRoutes())
}

// Route list functions - organize routes by domain/functionality
func ListPublicRoutes() []automatic_http.ApiRoute {
	return []automatic_http.ApiRoute{
		{
			Path:    "/status",
			Method:  http.MethodGet,
			Handler: statusHandler,
		},
	}
}

func ListUserRoutes() []automatic_http.ApiRoute {
	return []automatic_http.ApiRoute{
		{
			Path:    "/user/profile",
			Method:  http.MethodGet,
			Handler: getUserProfileHandler,
		},
		{
			Path:    "/user/profile",
			Method:  http.MethodPut,
			Handler: updateUserProfileHandler,
		},
	}
}

func ListAdminRoutes() []automatic_http.ApiRoute {
	return []automatic_http.ApiRoute{
		{
			Path:    "/users",
			Method:  http.MethodGet,
			Handler: listUsersHandler,
		},
		{
			Path:    "/users/:id",
			Method:  http.MethodDelete,
			Handler: deleteUserHandler,
		},
	}
}

func ListAgentRoutes() []automatic_http.ApiRoute {
	return []automatic_http.ApiRoute{
		{
			Path:    "/",
			Method:  http.MethodGet,
			Handler: listAgentsHandler,
		},
		{
			Path:    "/",
			Method:  http.MethodPost,
			Handler: createAgentHandler,
		},
		{
			Path:    "/:id",
			Method:  http.MethodGet,
			Handler: getAgentHandler,
		},
		{
			Path:    "/:id",
			Method:  http.MethodPut,
			Handler: updateAgentHandler,
		},
		{
			Path:    "/:id",
			Method:  http.MethodDelete,
			Handler: deleteAgentHandler,
		},
	}
}

// Handler functions
func statusHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"status": "ok", "version": "1.0.0"}, nil
}

func getUserProfileHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "get user profile"}, nil
}

func updateUserProfileHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "update user profile"}, nil
}

func listUsersHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "list users"}, nil
}

func deleteUserHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "delete user"}, nil
}

// Agent handlers
func listAgentsHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "list agents"}, nil
}

func createAgentHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "create agent"}, nil
}

func getAgentHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "get agent"}, nil
}

func updateAgentHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "update agent"}, nil
}

func deleteAgentHandler(c echo.Context) (any, errors.ApiError) {
	return map[string]string{"message": "delete agent"}, nil
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
