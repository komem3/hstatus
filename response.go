package hstatus

import (
	"encoding/json"
	"net/http"
)

type HTTPResp interface {
	Code() int
	WriteBody(http.ResponseWriter) error
	httpPrivate()
}

type (
	responseBase struct{}

	responseOKBase struct {
		*responseBase
	}
	responseCreatedBase struct {
		*responseBase
	}

	responseJSONBase struct {
		body interface{}
	}
	responseTextBase struct {
		text []byte
	}

	responseJSONOK struct {
		*responseOKBase
		*responseJSONBase
	}
	responseTextOK struct {
		*responseOKBase
		*responseTextBase
	}

	responseJSONCreated struct {
		*responseCreatedBase
		*responseJSONBase
	}
)

func (r *responseBase) httpPrivate() {}

func (r *responseOKBase) Code() int {
	return http.StatusOK
}

func (r *responseCreatedBase) Code() int {
	return http.StatusCreated
}

func (r *responseJSONBase) WriteBody(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(r.body)
}

func (r *responseTextBase) WriteBody(w http.ResponseWriter) error {
	_, err := w.Write(r.text)
	return err
}

// ResponseJSONOK response json and status ok.
func ResponseJSONOK(body interface{}) HTTPResp {
	return &responseJSONOK{responseJSONBase: &responseJSONBase{body: body}}
}

// ResponseTextOK response text and status ok.
func ResponseTextOK(text []byte) HTTPResp {
	return &responseTextOK{responseTextBase: &responseTextBase{text: text}}
}

// ResponseJSONCreated response json and status created.
func ResponseJSONCreated(body interface{}) HTTPResp {
	return &responseJSONCreated{responseJSONBase: &responseJSONBase{body: body}}
}
