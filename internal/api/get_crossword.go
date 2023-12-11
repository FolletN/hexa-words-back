package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	QueryParameterWidthKey  = "width"
	QueryParameterHeightKey = "height"
	defaultWidth            = 12
	defaultHeight           = 20
	minDimension            = 2
	maxDimension            = 50
)

func (s Server) HandleGetCrosswordGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received /crossword/grid command")

	width, height, err := getGridDimensions(
		r.URL.Query().Get(QueryParameterWidthKey),
		r.URL.Query().Get(QueryParameterHeightKey),
	)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	grid, err := s.crosswordGenerator.GenerateGrid(width, height)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
