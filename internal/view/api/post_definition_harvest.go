package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PostDefinitionHarvestParameters struct {
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty"`
}

func (s Server) HandlePostDefinitionHarvest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received /definition/harvest command")
	params := &PostDefinitionHarvestParameters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		fmt.Printf("failed to parse body : %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var startDate time.Time
	if params.StartDate != nil {
		startDate = *params.StartDate
	} else {
		startDate = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	var endDate time.Time
	if params.EndDate != nil {
		endDate = *params.EndDate
	} else {
		endDate = time.Now()
	}

	if err := s.definitionController.HarvestDefinitionsBetweenDates(r.Context(), startDate, endDate); err != nil {
		fmt.Printf("failed to harvest dates : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("/definition/harvest done")

	w.WriteHeader(http.StatusOK)
}
