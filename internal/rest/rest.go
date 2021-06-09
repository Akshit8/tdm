// Package rest deines handlers, controllers and middlewares for REST api.
package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Akshit8/tdm/internal"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func renderErrorResponse(w http.ResponseWriter, msg string, err error) {
	resp := ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  msg,
	}
	status := http.StatusInternalServerError

	var ierr *internal.Error
	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
	} else {
		switch ierr.Code() {
		case internal.ErrorCodeNotFound:
			resp.Status = http.StatusNotFound
			status = http.StatusNotFound
		case internal.ErrorCodeInvalidArgument:
			resp.Status = http.StatusBadRequest
			status = http.StatusBadRequest
		}
	}

	log.Printf("Error: %v\n", err)

	renderResponse(w, resp, status)
}

func renderResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	_, err = w.Write(content)
	if err != nil {
		log.Println("Couldn't write response", err)
	}
}
