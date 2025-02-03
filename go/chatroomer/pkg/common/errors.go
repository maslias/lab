package common

import (
	"errors"
	"net/http"
)

var (
	ErrInternalFailure = errors.New("internal failure")
	ErrBadRequest      = errors.New("bad request")
	ErrNotFound        = errors.New("not found")
	ErrAuth            = errors.New("unauthorized")
)

type Error struct {
	genericErr     error
	applicationErr error
}

func NewError(genericErr error, applicationErr error) *Error {
	return &Error{
		genericErr:     genericErr,
		applicationErr: applicationErr,
	}
}

func (e *Error) GenericErr() error {
	return e.genericErr
}

func (e *Error) ApplicationErr() error {
	return e.applicationErr
}

func (e *Error) Error() string {
	return errors.Join(e.genericErr, e.applicationErr).Error()
}

func GetErrorHttpStatus(err error) int {
	switch {
	case errors.As(ErrInternalFailure, err):
		return http.StatusInternalServerError
	case errors.As(ErrBadRequest, err):
		return http.StatusBadRequest
	case errors.As(ErrNotFound, err):
		return http.StatusNotFound
	case errors.As(ErrAuth, err):
		return http.StatusNonAuthoritativeInfo
	default:
		return 0
	}
}
