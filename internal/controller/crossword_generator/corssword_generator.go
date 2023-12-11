package crossword_generator

import "hexacrosswords/internal/db"

type CrosswordGenerator struct {
	DefinitionHandler db.DefinitionHandler
}

type Grid struct {
}

type Definition struct {
}

func (c CrosswordGenerator) GenerateGrid(width, height int) (Grid, error) {
	//
	return Grid{}, nil
}
