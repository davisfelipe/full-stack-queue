package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/davisfelipe/full-stack-queue/handlers"
	"github.com/davisfelipe/full-stack-queue/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

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

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.TestHandler(s)).Methods(http.MethodGet)
}
