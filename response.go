package hstatus

import "net/http"

type HTTPResponse interface {
	Body() interface{}
	Code() int
}

type (
	responseBase struct {
		body interface{}
	}

	responseOK struct {
		responseBase
	}
	responseCreate struct {
		responseBase
	}
)

func (r *responseBase) Body() interface{} {
	return r.body
}

func ResponseOK(body interface{}) HTTPResponse {
	return &responseOK{responseBase{body}}
}

func (r *responseOK) Code() int {
	return http.StatusOK
}

func ResponseCreate(body interface{}) HTTPResponse {
	return &responseCreate{responseBase{body}}
}

func (r *responseCreate) Code() int {
	return http.StatusCreated
}
