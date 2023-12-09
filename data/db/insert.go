package db

import (
	"cross-words-harverter/interprete"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type WordDefinition struct {
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
	return db.DB.Model((*WordDefinition)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}

func (db DBHandler) InsertWordDefinition(solution interprete.Solution) error {
	err := db.createWordDefinitionSchema()
	if err != nil {
		return err
	}

	wordDefinition := &WordDefinition{
		Definition: solution.Definition,
		Word:       solution.Word,
		Strength:   solution.Strength,
		Length:     len(solution.Word),
	}
	if len(solution.Word) >= 1 {
		wordDefinition.Char1 = string(solution.Word[0])
	}
	if len(solution.Word) >= 2 {
		wordDefinition.Char2 = string(solution.Word[1])
	}
	if len(solution.Word) >= 3 {
		wordDefinition.Char3 = string(solution.Word[2])
	}
	if len(solution.Word) >= 4 {
		wordDefinition.Char4 = string(solution.Word[3])
	}
	if len(solution.Word) >= 5 {
		wordDefinition.Char5 = string(solution.Word[4])
	}
	if len(solution.Word) >= 6 {
		wordDefinition.Char6 = string(solution.Word[5])
	}
	if len(solution.Word) >= 7 {
		wordDefinition.Char7 = string(solution.Word[6])
	}
	if len(solution.Word) >= 8 {
		wordDefinition.Char8 = string(solution.Word[7])
	}
	if len(solution.Word) >= 9 {
		wordDefinition.Char9 = string(solution.Word[8])
	}
	if len(solution.Word) >= 10 {
		wordDefinition.Char10 = string(solution.Word[9])
	}
	if len(solution.Word) >= 11 {
		wordDefinition.Char11 = string(solution.Word[10])
	}
	if len(solution.Word) >= 12 {
		wordDefinition.Char12 = string(solution.Word[11])
	}
	if len(solution.Word) >= 13 {
		wordDefinition.Char13 = string(solution.Word[12])
	}
	if len(solution.Word) >= 14 {
		wordDefinition.Char14 = string(solution.Word[13])
	}
	if len(solution.Word) >= 15 {
		wordDefinition.Char15 = string(solution.Word[14])
	}

	if len(solution.Word) >= 16 {
		fmt.Printf("ERROR: word longer than 16 characters\n")
	}

	_, err = db.DB.Model(wordDefinition).OnConflict("(definition, word, strength) DO NOTHING").Insert()
	return err
}
