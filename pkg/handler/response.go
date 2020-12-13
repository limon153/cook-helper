package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type responseError struct {
	Message    string `json:"message"`
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

func newResponseError(rw http.ResponseWriter, status int, message string) {
	log.Print(message)
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")

	res, err := json.Marshal(responseError{
		message,
		status,
		http.StatusText(status),
	})

	if err != nil {
		log.Printf("error forming json error response: %s", err.Error())
		rw.Write([]byte("could not create json error message"))
	}

	rw.Write(res)
}

func newResponseOk(rw http.ResponseWriter, body map[string]interface{}) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")

	res, err := json.Marshal(body)
	if err != nil {
		newResponseError(rw, http.StatusInternalServerError, err.Error())
	}

	rw.Write(res)
}
