package database

import (
	"fmt"

	"cross-words/internal/controller/definition_collector"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Definition struct {
	Definition string `pg:"definition,unique:line" `
	Word       string `pg:"word,unique:line"`
	Strength   int    `pg:"strength,unique:line"`
	Length     int    `pg:"length"`
	Char1      string `pg:"1"`
	Char2      string `pg:"2"`
	Char3      string `pg:"3"`
	Char4      string `pg:"4"`
	Char5      string `pg:"5"`
	Char6      string `pg:"6"`
	Char7      string `pg:"7"`
	Char8      string `pg:"8"`
	Char9      string `pg:"9"`
	Char10     string `pg:"10"`
	Char11     string `pg:"11"`
	Char12     string `pg:"12"`
	Char13     string `pg:"13"`
	Char14     string `pg:"14"`
	Char15     string `pg:"15"`
}

const (
	database = "postgres"
	user     = "crossword"
	password = "crossword"
)

func NewDB() DBHandler {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")
	options := &pg.Options{
		User:     user,
		Password: password,
		Addr:     address,
		Database: database,
	}

	DB := pg.Connect(options)
	return DBHandler{
		DB: DB,
	}
}

func (db DBHandler) createWordDefinitionSchema() error {
	return db.DB.Model((*Definition)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}

func (db DBHandler) InsertWordDefinition(definitions definition_collector.Definition) error {
	err := db.createWordDefinitionSchema()
	if err != nil {
		return err
	}

	wordDefinition := &Definition{
		Definition: definitions.Definition,
		Word:       definitions.Word,
		Strength:   definitions.Strength,
		Length:     len(definitions.Word),
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

	_, err = db.DB.Model(wordDefinition).OnConflict("(definition, word, strength) DO NOTHING").Insert()
	return err
}
