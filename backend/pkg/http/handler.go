package http

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/notzree/automaticv2/v2/pkg/errors"

	"github.com/labstack/echo/v4"
)

// HandlerFunc type accepts an echo context and returns a response object and error
// The response object can be a struct, slice, map, string, or any other in-built type that can be marshalled into JSON
type HandlerFunc func(c echo.Context) (any, errors.ApiError)

// Wrap is a simple wrapper that accepts a HandlerFunc and returns an echo.HandlerFunc
// This allows HandlerFunc to behave as normal Go functions that return 'any' object and error, and all the
// echo specific response handling is abstracted into this wrapper function
// Note that the error_handler middleware will construct the appropriate response based on the echo context errors
func Wrap(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r, err := h(c)
		if err != nil {
			handleApiError(c, err)
			return err
		}
		return c.JSON(http.StatusOK, r)
	}
}

// WrapWithoutStatus is a simple wrapper that accepts a HandlerFunc and returns an echo.HandlerFunc
// There will be no JSON serialization for file responses as this corrupts the file content
// When serialized, the response is appended to the bytes of the file. This causes excel to report an error
func WrapWithoutStatus(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := h(c)
		if err != nil {
			handleApiError(c, err)
			return err
		}
		return nil
	}
}

// handleApiError writes the appropriate error response to echo response writer and emits a log
func handleApiError(c echo.Context, apiErr errors.ApiError) {
	statusCode := apiErr.StatusCode()
	errorMsg := apiErr.DisplayErrorMessage()
	slog.ErrorContext(c.Request().Context(),
		fmt.Sprintf("%s. Error: %v", errorMsg, apiErr.Unwrap()),
		slog.Any("error", apiErr),
	)
	c.JSON(statusCode, map[string]string{"error": errorMsg})
}
