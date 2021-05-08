package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type ConfigServer struct {
	Address string `yaml:"address"`
}

type Server struct {
	lg     *logrus.Logger
	config *ConfigServer
}

func NewServer(config *ConfigServer, lg *logrus.Logger) *Server {
	return &Server{
		lg:     lg,
		config: config,
	}
}

func (server *Server) Start() error {
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Get("/", server.HandleGetIndex)
	})
	return http.ListenAndServe(server.config.Address, router)
}
