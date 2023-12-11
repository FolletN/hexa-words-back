package harvester

import (
	"context"
	"fmt"
	"sync"
	"time"

	"hexacrosswords/internal/controller/harvester/collector"
	"hexacrosswords/internal/db"

	"golang.org/x/sync/errgroup"
)

type Harvester struct {
	Collectors        []collector.Collector
	DefinitionHandler db.DefinitionHandler
}

func (h Harvester) HarvestDefinitionsBetweenDates(ctx context.Context, startDate time.Time, endDate time.Time) error {
	fmt.Printf("harvesting dates [%s, %s]\n", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	// chan used to send definition to be stored
	definitions := []db.Definition{}
	mu := sync.Mutex{}

	var g errgroup.Group
	harvestingDate := startDate
	for harvestingDate.Before(endDate) {
		for _, definitionCollector := range h.Collectors {
			func(copiedHarvestingDate time.Time) {
				g.Go(func() error {
					newDefinitions, err := definitionCollector.GetDefinitions(ctx, copiedHarvestingDate)
					if err != nil {
						return fmt.Errorf("failed to get definitions at date %s : %s", copiedHarvestingDate, err.Error())
					}
					mu.Lock()
					definitions = append(definitions, newDefinitions...)
					mu.Unlock()
					return nil
				})
			}(harvestingDate)
		}
		harvestingDate = harvestingDate.AddDate(0, 0, 1)
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("error while harvesting definition : %s", err.Error())
		return err
	}

	if err := h.DefinitionHandler.StoreDefinitions(ctx, definitions); err != nil {
		return fmt.Errorf("failed to store definitions : %s", err)
	}

	return nil
}
