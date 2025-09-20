package errors

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
)

// WithApiError wraps around echo's c.Error and handles the response
// Additionally, we enforce passing an apiError instead of error so that we restrict this usage to handler functions
// This function writes the appropriate error response and logs the error
func WithApiError(c echo.Context, err ApiError) error {
	statusCode := err.StatusCode()
	errorMsg := err.DisplayErrorMessage()
	slog.ErrorContext(c.Request().Context(),
		fmt.Sprintf("%s. Error: %v", errorMsg, err.Unwrap()),
		slog.Any("error", err),
	)
	return c.JSON(statusCode, map[string]string{"error": errorMsg})
}
