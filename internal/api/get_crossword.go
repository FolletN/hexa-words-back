package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	QueryParameterWidthKey    = "width"
	QueryParameterHeightKey   = "height"
	QueryParameterStrengthKey = "strength"
	defaultWidth              = 12
	defaultHeight             = 20
	minDimension              = 2
	maxDimension              = 50
)

func (s Server) HandleGetCrosswordGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received /crossword/grid command")

	query := r.URL.Query()
	width, height, err := getGridDimensions(
		query.Get(QueryParameterWidthKey),
		query.Get(QueryParameterHeightKey),
	)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	strength, err := getGridStrength(
		query.Get(QueryParameterStrengthKey),
	)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	grid, err := s.crosswordGenerator.GenerateGrid(r.Context(), strength, width, height)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(grid)
	w.WriteHeader(http.StatusOK)
	fmt.Println("/crossword/grid done")
}

func getGridDimensions(widthStr, heightStr string) (int, int, error) {
	width := defaultWidth
	if widthStr != "" {
		var err error
		width, err = strconv.Atoi(widthStr)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse width : %w", err)
		}
	}

	height := defaultHeight
	if heightStr != "" {
		var err error
		height, err = strconv.Atoi(heightStr)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse height : %w", err)
		}
	}

	if width < minDimension || height < minDimension || width > maxDimension || height > maxDimension {
		return 0, 0, fmt.Errorf("dimensions must be between %v and %v", minDimension, maxDimension)
	}

	return width, height, nil
}

func getGridStrength(strengthStr string) (int, error) {
	strength := 2
	if strengthStr != "" {
		var err error
		strength, err = strconv.Atoi(strengthStr)
		if err != nil {
			return 0, fmt.Errorf("failed to parse strength : %w", err)
		}
	}

	if strength > 4 || strength < 1 {
		return 0, fmt.Errorf("strength out or range : %v", strength)
	}
	return strength, nil
}
