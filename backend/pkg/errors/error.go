package errors

import (
	stderr "errors"

	"github.com/pkg/errors"
)

// The functions in this file provide a simple wrapper around functionalities in the pkg/errors and std errors packages
// This avoids having to import multiple errors packages across the backend codebase
// Is, As and Unwrap functions are also exposed by the pkg/errors package, but they are simply wrappers around the
// functions exposed by the standard errors package, so we avoid the additional function call by wrapping those functions ourselves

func New(text string) error {
	return errors.New(text)
}

func Wrap(err error, msg string) error {
	// handle edge cases where we try to wrap a nil error with a non-empty msg
	if err == nil && msg != "" {
		return errors.New(msg)
	}
	return errors.Wrap(err, msg)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// Standard errors package

func Is(err error, target error) bool {
	return stderr.Is(err, target)
}

func As(err error, target interface{}) bool {
	return stderr.As(err, target)
}

func Unwrap(err error) error {
	return stderr.Unwrap(err)
}

func Join(errs ...error) error {
	return stderr.Join(errs...)
}
