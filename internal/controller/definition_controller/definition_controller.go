package definition_controller

import (
	"cross-words/internal/controller/definition_collector"
	"cross-words/internal/model/database"
)

type DefinitionController struct {
	DefinitionCollector definition_collector.DefinitionCollector
	DataHandler         database.DBHandler
}

func NewDefinitionController() DefinitionController {
	return DefinitionController{
		DefinitionCollector: definition_collector.NewDefinitionCollector(),
		DataHandler:         database.NewDB(),
	}
}
