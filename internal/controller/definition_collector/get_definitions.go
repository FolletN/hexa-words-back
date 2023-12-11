package definition_collector

import (
	"context"
	"cross-words/external/vingt_minutes"
	"fmt"
	"time"
)

func (d DefinitionCollector) GetDefinitions(ctx context.Context, date time.Time) ([]Definition, error) {
	formatedDate := date.Format("020106")
	fmt.Printf("collection definition of date %v\n", formatedDate)

	data, err := vingt_minutes.GetGameData(ctx, formatedDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get game data")
	}

	solutions, err := d.ParseDefinitions(data)
	if err != nil {
		return nil, fmt.Errorf("failed to interprete data : %s", err.Error())
	}
	if len(solutions) == 0 {
		return nil, fmt.Errorf("failed to interprete data : no data")
	}

	return solutions, nil
}
