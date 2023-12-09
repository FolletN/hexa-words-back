package db

import (
	"context"
	"cross-words-harverter/interprete"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBHandler struct {
	DB *pg.DB
}

func (db DBHandler) StoreSolutions(ctx context.Context, solutions []interprete.Solution) error {
	for _, solution := range solutions {
		if err := db.InsertWordDefinition(solution); err != nil {
			fmt.Printf("Error while storing data : %v\n", err.Error())
			return fmt.Errorf("failed to store data")
		}
	}
	return nil
}
