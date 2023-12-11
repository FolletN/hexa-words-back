package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
)

type DefinitionHandler struct {
	DB *bun.DB
}

type Definition struct {
	Statement string
	Word      string
	Strength  int
}

func (handler DefinitionHandler) StoreDefinitions(ctx context.Context, definitions []Definition) error {
	err := handler.DB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for _, definition := range definitions {
			if err := storeDefinition(ctx, tx, definition); err != nil {
				fmt.Printf("failed to store definition : %s", err)
				return err
			}
		}
		return nil
	})
	return err
}

type dbDefinition struct {
	bun.BaseModel `bun:"table:definition,alias:def"`

	ID int64 `bun:"id,pk,autoincrement"`

	Statement string `bun:"statement,unique:line" `
	Word      string `bun:"word,unique:line"`
	Strength  int    `bun:"strength,unique:line"`
	Length    int    `bun:"length,unique:line"`
	Char1     string `bun:"1"`
	Char2     string `bun:"2"`
	Char3     string `bun:"3"`
	Char4     string `bun:"4"`
	Char5     string `bun:"5"`
	Char6     string `bun:"6"`
	Char7     string `bun:"7"`
	Char8     string `bun:"8"`
	Char9     string `bun:"9"`
	Char10    string `bun:"10"`
	Char11    string `bun:"11"`
	Char12    string `bun:"12"`
	Char13    string `bun:"13"`
	Char14    string `bun:"14"`
	Char15    string `bun:"15"`
}

func createWordDefinitionSchema(ctx context.Context, db bun.IDB) error {
	_, err := db.NewCreateTable().Model((*dbDefinition)(nil)).IfNotExists().Exec(ctx)
	return err
}

func storeDefinition(ctx context.Context, db bun.IDB, definitions Definition) error {
	if err := createWordDefinitionSchema(ctx, db); err != nil {
		return err
	}

	wordDefinition := &dbDefinition{
		Statement: definitions.Statement,
		Word:      definitions.Word,
		Strength:  definitions.Strength,
		Length:    len(definitions.Word),
	}
	if len(definitions.Word) >= 1 {
		wordDefinition.Char1 = string(definitions.Word[0])
	}
	if len(definitions.Word) >= 2 {
		wordDefinition.Char2 = string(definitions.Word[1])
	}
	if len(definitions.Word) >= 3 {
		wordDefinition.Char3 = string(definitions.Word[2])
	}
	if len(definitions.Word) >= 4 {
		wordDefinition.Char4 = string(definitions.Word[3])
	}
	if len(definitions.Word) >= 5 {
		wordDefinition.Char5 = string(definitions.Word[4])
	}
	if len(definitions.Word) >= 6 {
		wordDefinition.Char6 = string(definitions.Word[5])
	}
	if len(definitions.Word) >= 7 {
		wordDefinition.Char7 = string(definitions.Word[6])
	}
	if len(definitions.Word) >= 8 {
		wordDefinition.Char8 = string(definitions.Word[7])
	}
	if len(definitions.Word) >= 9 {
		wordDefinition.Char9 = string(definitions.Word[8])
	}
	if len(definitions.Word) >= 10 {
		wordDefinition.Char10 = string(definitions.Word[9])
	}
	if len(definitions.Word) >= 11 {
		wordDefinition.Char11 = string(definitions.Word[10])
	}
	if len(definitions.Word) >= 12 {
		wordDefinition.Char12 = string(definitions.Word[11])
	}
	if len(definitions.Word) >= 13 {
		wordDefinition.Char13 = string(definitions.Word[12])
	}
	if len(definitions.Word) >= 14 {
		wordDefinition.Char14 = string(definitions.Word[13])
	}
	if len(definitions.Word) >= 15 {
		wordDefinition.Char15 = string(definitions.Word[14])
	}

	if len(definitions.Word) >= 16 {
		fmt.Printf("ERROR: word longer than 16 characters\n")
	}

	_, err := db.NewInsert().Model(wordDefinition).On("CONFLICT (statement, word, strength, length) DO NOTHING").Exec(ctx)
	return err
}

type SearchDefinitionParameters struct {
}

func (db DefinitionHandler) SearchDefinition(ctx context.Context, searchParameters SearchDefinitionParameters) ([]Definition, error) {
	return nil, nil
}
