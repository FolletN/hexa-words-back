package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PostHarvestParameters struct {
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty"`
}

func (s Server) HandlePostHarvest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received /harvest command")
	params := &PostHarvestParameters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		fmt.Printf("failed to parse body : %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	if params.StartDate != nil {
		startDate = *params.StartDate
	}

	endDate := time.Now()
	if params.EndDate != nil {
		endDate = *params.EndDate
	}

	if err := s.definitionHarvester.HarvestDefinitionsBetweenDates(r.Context(), startDate, endDate); err != nil {
		w.Write([]byte(fmt.Sprintf("failed to harvest dates : %s", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("/harvest done")
	w.WriteHeader(http.StatusOK)
}
