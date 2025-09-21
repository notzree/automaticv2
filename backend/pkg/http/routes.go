package http

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ApiRoute represents a single API route
type ApiRoute struct {
	Path    string
	Method  string
	Handler HandlerFunc
}

// PermissionFunc represents a permission check function
type PermissionFunc func(c echo.Context) error

// RegisterHandlers registers multiple API routes under a specific Echo group
func RegisterHandlers(group *echo.Group, routes []ApiRoute) {
	for _, route := range routes {
		echoHandler := Wrap(route.Handler)
		switch route.Method {
		case http.MethodGet:
			group.GET(route.Path, echoHandler)
		case http.MethodPost:
			group.POST(route.Path, echoHandler)
		case http.MethodPut:
			group.PUT(route.Path, echoHandler)
		case http.MethodDelete:
			group.DELETE(route.Path, echoHandler)
		case http.MethodPatch:
			group.PATCH(route.Path, echoHandler)
		case http.MethodOptions:
			group.OPTIONS(route.Path, echoHandler)
		default:
			// Log error but continue with other routes
			slog.Error("invalid method", "method", route.Method)
			continue
		}
	}
}

// Example permission functions (as middleware)
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement authentication check
		return next(c)
	}
}

func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement admin permission check
		return next(c)
	}
}

// Example middleware functions
func LoggingMiddleware() echo.MiddlewareFunc {
	// TODO: Implement logging middleware
	return middleware.Logger()
}

func CORSMiddleware() echo.MiddlewareFunc {
	return middleware.CORS()
}
