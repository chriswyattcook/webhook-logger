package main

import (
	"encoding/json"
	"net/http"
)

type applicationData struct {
	Status string `json:"status"`
}

func health(w http.ResponseWriter, r *http.Request) {
	data := applicationData{
		Status: "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
