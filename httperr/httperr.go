package httperr

import (
	"net/http"
)

type HttpErr interface {
	Code() int
	Error() string
}

type NotFound struct {
	code    uint
	message string
	err     error
}

func NewNotFound(message string, err error) *NotFound {
	return &NotFound{code: http.StatusNotFound, message: message, err: err}
}

func (e *NotFound) Error() string {
	return "Not Found: " + e.message
}

func (e *NotFound) Code() int {
	return int(e.code)
}

type InternalServer struct {
	code    uint
	message string
	err     error
}

func NewInternalServer(message string, err error) *InternalServer {
	return &InternalServer{code: http.StatusInternalServerError, message: message, err: err}
}

func (e *InternalServer) Error() string {
	return "Internal server error: " + e.message
}

func (e *InternalServer) Code() int {
	return int(e.code)
}
