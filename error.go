package hstatus

import "net/http"

type ErrorResp interface {
	Err() error
	Code() int
}

type (
	errBase struct {
		error
	}

	errBadRequest struct {
		errBase
	}
	errNotFound struct {
		errBase
	}
	errInternalServerError struct {
		errBase
	}
)

func (e *errBase) Err() error {
	return e.error
}

func ErrBadRequest(err error) ErrorResp {
	return &errBadRequest{errBase{err}}
}

func (e *errBadRequest) Code() int {
	return http.StatusBadRequest
}

func ErrNotFound(err error) ErrorResp {
	return &errNotFound{errBase{err}}
}

func (e *errNotFound) Code() int {
	return http.StatusNotFound
}

func ErrInternalServerError(err error) ErrorResp {
	return &errInternalServerError{errBase{err}}
}

func (e *errInternalServerError) Code() int {
	return http.StatusInternalServerError
}
