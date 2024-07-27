package crossword_generator

import (
	"context"
	"hexacrosswords/internal/db"
	"hexacrosswords/internal/utils/convert"
	"hexacrosswords/internal/utils/math"
)

type CrosswordGenerator struct {
	DefinitionHandler db.DefinitionHandler
}

type Grid struct {
	Squares [][]Square `json:"squares"`
}

type Coordinates struct {
	X int
	Y int
}

type Direction string

const (
	X Direction = "X"
	Y Direction = "Y"
)

func (g *Grid) AddWordToGrid(definitionCoordinates Coordinates, firstLetterCoordinates Coordinates, direction Direction, definition db.Definition) {

}

type Square struct {
}

type Definition struct {
}

func (c CrosswordGenerator) GenerateGrid(ctx context.Context, strength, width, height int) (*Grid, error) {
	// Initialize grid - first dimention is X (width), second dimension is Y (height)
	squares := make([][]Square, width)
	for i := range squares {
		squares[i] = make([]Square, height)
	}

	grid := &Grid{
		Squares: squares,
	}

	// (0;0) is the starting point, need 2 definitions
	// the first one being for vertical definition starting at (1;0)
	// the seconf one being for horizontal definition starting at (0;1)
	// we want both definition being as long as possible, between 10 and 15 letters long (maxed by remaning size of max width)

	// first definition
	firstVerticalDefinition, err := c.DefinitionHandler.SearchDefinition(ctx, db.SearchDefinitionParameters{
		MaxStrength: &strength,
		OrderBy: []string{
			"strength desc",
		},
		OrderExpr: convert.StringToPtr("random()"),
		MaxLength: convert.IntToPtr(math.MaxInt(15, height-1)),
		MinLength: convert.IntToPtr(math.MinInt(10, height-1)),
	})
	if err != nil {
		return nil, err
	}

	grid.AddWordToGrid(
		Coordinates{
			X: 0,
			Y: 0,
		},
		Coordinates{
			X: 1,
			Y: 0,
		},
		Y,
		firstVerticalDefinition,
	)

	return grid, nil
}
