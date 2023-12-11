package database

import (
	"context"
	"fmt"

	"cross-words/internal/controller/definition_collector"

	"github.com/go-pg/pg/v10"
)

type DBHandler struct {
	DB *pg.DB
}

func (db DBHandler) StoreDefinitions(ctx context.Context, definitions []definition_collector.Definition) error {
	for _, definition := range definitions {
		if err := db.InsertWordDefinition(definition); err != nil {
			fmt.Printf("Error while storing data : %v\n", err.Error())
			return fmt.Errorf("failed to store data")
		}
	}
	return nil
}
