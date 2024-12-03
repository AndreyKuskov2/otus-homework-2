package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func health(w http.ResponseWriter, r *http.Request) {
	healthResponse := HealthResponse{
		Status: "ok",
	}

	resp, err := json.Marshal(healthResponse)
	if err != nil {
		http.Error(w, "Failed to convert response!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	log.Info("Start server on 8000 port!")
	r := chi.NewRouter()

	r.Get("/health/", health)

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Errorf("Error starting server: %s", err.Error())
		return
	}
}
