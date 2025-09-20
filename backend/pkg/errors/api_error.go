package errors

import (
	"fmt"
	"net/http"
)

// ApiError interface implements the built-in error interface
// It represents errors that must be returned by api handlers
type ApiError interface {
	Error() string
	StatusCode() int
	DisplayMessage() string
	DisplayErrorMessage() string
	Unwrap() error
	APIError() (int, string)
}

// apiError is an error type used for returning API errors from API handlers
// this error type implements the ApiError interface and contains additional metadata
// err field is used to access the underlying error when emitting logs, this is never returned to the client
// statusCode, statusText, message are returned to the client
type apiError struct {
	err        error
	statusCode int
	statusText string
	message    string
}

func (e *apiError) Error() string {
	msg := fmt.Sprintf("API error %d: %s.", e.statusCode, e.statusText)
	if e.message != "" {
		msg = fmt.Sprintf(" %s. %s", msg, e.message)
	}
	if e.err != nil {
		msg = fmt.Sprintf("%s. Error: %s", msg, e.err.Error())
	}
	return msg
}

func (e *apiError) Unwrap() error {
	return e.err
}

func (e *apiError) StatusCode() int {
	return e.statusCode
}

func (e *apiError) DisplayMessage() string {
	return e.message
}

func (e *apiError) APIError() (int, string) {
	return e.statusCode, e.statusText
}

func (e *apiError) DisplayErrorMessage() string {
	errorMsg := http.StatusText(e.StatusCode())
	if e.DisplayMessage() != "" {
		errorMsg = fmt.Sprintf("%s. %s", errorMsg, e.DisplayMessage())
	}
	return errorMsg
}

// WithDisplayMessage attaches a display message that is returned to the client,
// so this should not contain sensitive error specific information
func (e *apiError) WithDisplayMessage(message string) *apiError {
	e.message = message
	return e
}

// WithError is only used for logging
// To display a message to the client, use WithDisplayMessage
func (e *apiError) WithError(err error) *apiError {
	e.err = err
	return e
}

func NewAPIError(statusCode int) *apiError {
	// default to internal server error if status code is out of range of standard HTTP status codes
	if statusCode < 100 || statusCode >= 600 {
		statusCode = http.StatusInternalServerError
	}
	return &apiError{statusCode: statusCode, statusText: http.StatusText(statusCode)}
}

func APIErrorOrNewError(err error, msg string) error {
	if _, ok := err.(*apiError); ok {
		return err
	}
	return Wrap(err, msg)
}

func APIErrorOrNewAPIError(err error, statusCode int, msg string) ApiError {
	if apiErr, ok := err.(ApiError); ok {
		return apiErr
	}
	return NewAPIError(statusCode).WithDisplayMessage(msg).WithError(Wrap(err, msg))
}
