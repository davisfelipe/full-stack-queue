package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/davisfelipe/full-stack-queue/handlers"
	"github.com/davisfelipe/full-stack-queue/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type DefaultErrorHandler struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(BindRoutes)
}

func CustomeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + " " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func CustomNotFoundError(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		log.Println(r.Method + " " + r.RequestURI)
		json.NewEncoder(w).Encode(DefaultErrorHandler{
			Message: "Route not found",
			Status:  false,
		})
	})
}

func BindRoutes(s server.Server, r *mux.Router) {

	// Middleware
	r.Use(CustomeMiddleware)
	r.NotFoundHandler = CustomNotFoundError(r)

	// Routes
	r.HandleFunc("/", handlers.TestHandler(s)).Methods(http.MethodGet)
}
