package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/komem3/hstatus"
	"github.com/rs/xid"
)

func responseConv(f func(*http.Request) (hstatus.HTTPResponse, hstatus.ErrorResp)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, herr := f(r)
		if herr != nil {
			w.WriteHeader(herr.Code())
			w.Write([]byte(herr.Err().Error()))
			return
		}
		w.WriteHeader(resp.Code())
		if err := json.NewEncoder(w).Encode(resp.Body()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func sampleRouter(r chi.Router) {
	r.Get("/{id}", responseConv(get))
	r.Post("/", responseConv(create))
}

func get(r *http.Request) (hstatus.HTTPResponse, hstatus.ErrorResp) {
	id := chi.URLParam(r, "id")

	if id == "" {
		return nil, hstatus.ErrBadRequest(errors.New("id is empty"))
	}
	return hstatus.ResponseOK("id is " + id), nil
}

func create(r *http.Request) (hstatus.HTTPResponse, hstatus.ErrorResp) {
	type (
		reqPost struct {
			Content string `json:"content"`
		}
		respPost struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}
	)
	var req reqPost
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, hstatus.ErrBadRequest(fmt.Errorf("request parse: %w", err))
	}
	id := xid.New().String()

	return hstatus.ResponseCreate(respPost{ID: id, Content: req.Content}), nil
}
