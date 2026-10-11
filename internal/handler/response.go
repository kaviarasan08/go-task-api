package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct{
	Status int `json:"status"`
	Error string `json:"error"`
}

func writJSONEror(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Contetnt-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{Error: message, Status: status})

}