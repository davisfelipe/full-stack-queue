package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port string
}

type Broker struct {
	config *Config
	router mux.Router
}

type Server interface {
	Config() *Config
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	broker := &Broker{
		config: config,
		router: *mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = *mux.NewRouter()
	binder(b, &b.router)
	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, &b.router); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
