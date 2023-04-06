package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/davisfelipe/full-stack-queue/server"
)

type TestResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func TestHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(TestResponse{
			Message: "Welcome to testing",
			Status:  true,
		})
	}
}
