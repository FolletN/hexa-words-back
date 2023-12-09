package handler

import (
	"context"
	"cross-words-harverter/data"
	"cross-words-harverter/httpclient"
	"cross-words-harverter/interprete"
	"fmt"
	"time"
)

type Handler struct {
	DataHandler data.Data
}

func (h Handler) Handle(ctx context.Context) error {
	date := time.Now()
	for date.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)) {
		if err := h.HandleDate(ctx, date); err != nil {
			return err
		}
		date = date.AddDate(0, 0, -1)
	}
	return nil
}

func (h Handler) HandleDate(ctx context.Context, date time.Time) error {
	formatedDate := date.Format("020106")
	solutions, err := getSolutions(ctx, formatedDate)
	if err != nil {
		return fmt.Errorf("failed to retrieve solutions for date %s : %s", formatedDate, err.Error())
	}
	if err := h.DataHandler.StoreSolutions(ctx, solutions); err != nil {
		return fmt.Errorf("failed to process solutions for date %s : %s", formatedDate, err.Error())
	}

	return nil
}

func getSolutions(ctx context.Context, formatedDate string) ([]interprete.Solution, error) {

	fmt.Printf("Processing date %v\n", formatedDate)

	data, err := httpclient.GetData(ctx, formatedDate)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve http data")
	}

	solutions, err := interprete.NewInterpretor(data).Interprete()
	if err != nil {
		return nil, fmt.Errorf("failed to interprete data : %s", err.Error())
	}
	if len(solutions) == 0 {
		return nil, fmt.Errorf("failed to interprete data : no data")
	}

	return solutions, nil
}
