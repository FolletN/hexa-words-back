package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"cross-words/internal/controller/definition_controller"
)

type Server struct {
	config               ServerConfiguration
	mux                  *mux.Router
	definitionController definition_controller.DefinitionController
}

func NewClient(config ServerConfiguration) Server {
	mux := mux.NewRouter()
	server := Server{
		config:               config,
		mux:                  mux,
		definitionController: definition_controller.NewDefinitionController(),
	}

	mux.HandleFunc("/definition/harvest", server.HandlePostDefinitionHarvest).Methods("POST")

	return server
}

func (s Server) Serve(ctx context.Context) error {
	fmt.Println("Server is listening")
	return http.ListenAndServe(fmt.Sprintf(":%v", s.config.Port), s.mux)
}
