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
	Strength    *int
	MaxStrength *int
	MinStrength *int
	Length      *int
	MaxLength   *int
	MinLength   *int
	OrderBy     []string
	OrderExpr   *string
	Char1       *string
	Char2       *string
	Char3       *string
	Char4       *string
	Char5       *string
	Char6       *string
	Char7       *string
	Char8       *string
	Char9       *string
	Char10      *string
	Char11      *string
	Char12      *string
	Char13      *string
	Char14      *string
	Char15      *string
}

func (db DefinitionHandler) SearchDefinition(ctx context.Context, searchParameters SearchDefinitionParameters) (Definition, error) {
	definition := dbDefinition{}
	query := db.DB.NewSelect().Model(&definition)

	if searchParameters.Strength != nil {
		query.Where("strength = ?", searchParameters.Strength)
	}
	if searchParameters.MinStrength != nil {
		query.Where("strength >= ?", searchParameters.MinStrength)
	}
	if searchParameters.MaxStrength != nil {
		query.Where("strength <= ?", searchParameters.MaxStrength)
	}
	if searchParameters.Length != nil {
		query.Where("length = ?", searchParameters.Strength)
	}
	if searchParameters.MinLength != nil {
		query.Where("length >= ?", searchParameters.MinLength)
	}
	if searchParameters.MaxLength != nil {
		query.Where("length <= ?", searchParameters.MaxLength)
	}
	if searchParameters.Char1 != nil {
		query.Where("1 = ?", searchParameters.Char1)
	}
	if searchParameters.Char2 != nil {
		query.Where("2 = ?", searchParameters.Char2)
	}
	if searchParameters.Char3 != nil {
		query.Where("3 = ?", searchParameters.Char3)
	}
	if searchParameters.Char4 != nil {
		query.Where("4 = ?", searchParameters.Char4)
	}
	if searchParameters.Char5 != nil {
		query.Where("5 = ?", searchParameters.Char5)
	}
	if searchParameters.Char6 != nil {
		query.Where("6 = ?", searchParameters.Char6)
	}
	if searchParameters.Char7 != nil {
		query.Where("7 = ?", searchParameters.Char7)
	}
	if searchParameters.Char8 != nil {
		query.Where("8 = ?", searchParameters.Char8)
	}
	if searchParameters.Char9 != nil {
		query.Where("9 = ?", searchParameters.Char9)
	}
	if searchParameters.Char10 != nil {
		query.Where("10 = ?", searchParameters.Char10)
	}
	if searchParameters.Char11 != nil {
		query.Where("11 = ?", searchParameters.Char11)
	}
	if searchParameters.Char12 != nil {
		query.Where("12 = ?", searchParameters.Char12)
	}
	if searchParameters.Char13 != nil {
		query.Where("13 = ?", searchParameters.Char13)
	}
	if searchParameters.Char14 != nil {
		query.Where("14 = ?", searchParameters.Char14)
	}
	if searchParameters.Char15 != nil {
		query.Where("15 = ?", searchParameters.Char15)
	}
	if searchParameters.OrderBy != nil {
		query.Order(searchParameters.OrderBy...)
	}
	if searchParameters.OrderBy != nil {
		query.OrderExpr(*searchParameters.OrderExpr)
	}

	fmt.Println(query.String())

	err := query.Scan(ctx)
	if err != nil {
		return Definition{}, err
	}

	return Definition{
		Statement: definition.Statement,
		Word:      definition.Word,
		Strength:  definition.Strength,
	}, nil
}
