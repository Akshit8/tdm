// Package rest deines handlers, controllers and middlewares for REST api.
package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func renderErrorResponse(w http.ResponseWriter, msg string, status int) {
	renderResponse(w, errorResponse{Status: status, Error: msg}, status)
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
