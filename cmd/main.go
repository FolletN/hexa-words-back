package main

import (
	"context"
	"fmt"

	"hexacrosswords/internal/api"
	"hexacrosswords/internal/configuration"
	"hexacrosswords/internal/controller/crossword_generator"
	"hexacrosswords/internal/controller/harvester"
	"hexacrosswords/internal/controller/harvester/collector"
	"hexacrosswords/internal/db"
)

func main() {
	ctx := context.Background()

	fmt.Println("Initializing server")
	config, err := configuration.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	conn := db.NewDB(config.DB)
	if err := conn.Ping(); err != nil {
		panic(err)
	}

	definitionHandler := db.DefinitionHandler{
		DB: conn,
	}
	definitionHarvester := harvester.Harvester{
		Collectors: []collector.Collector{
			collector.GetCollector20Minutes(),
		},
		DefinitionHandler: definitionHandler,
	}
	crosswordGenerator := crossword_generator.CrosswordGenerator{
		DefinitionHandler: definitionHandler,
	}
	server := api.NewClient(
		config.Server,
		definitionHarvester,
		crosswordGenerator,
	)

	if err := server.Serve(ctx); err != nil {
		panic(err)
	}
}
