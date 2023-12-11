package api

import (
	"context"
	"fmt"
	"net/http"

	"hexacrosswords/internal/controller/crossword_generator"
	"hexacrosswords/internal/controller/harvester"

	"github.com/gorilla/mux"
)

type Configuration struct {
	Port int
}

type Server struct {
	config              Configuration
	mux                 *mux.Router
	definitionHarvester harvester.Harvester
	crosswordGenerator  crossword_generator.CrosswordGenerator
}

func NewClient(
	config Configuration,
	definitionHarvester harvester.Harvester,
	crosswordGenerator crossword_generator.CrosswordGenerator,
) Server {
	router := mux.NewRouter()

	server := Server{
		config:              config,
		mux:                 router,
		definitionHarvester: definitionHarvester,
		crosswordGenerator:  crosswordGenerator,
	}

	router.HandleFunc("/harvest", server.HandlePostHarvest).Methods(http.MethodPost)
	router.HandleFunc("/crossword/grid", server.HandleGetCrosswordGrid).Methods(http.MethodGet)

	return server
}

func (s Server) Serve(ctx context.Context) error {
	fmt.Println("Server is listening")
	return http.ListenAndServe(fmt.Sprintf(":%v", s.config.Port), s.mux)
}
