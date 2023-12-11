package definition_controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"cross-words/internal/controller/definition_collector"

	"golang.org/x/sync/errgroup"
)

func (d DefinitionController) HarvestDefinitionsBetweenDates(ctx context.Context, startDate time.Time, endDate time.Time) error {
	fmt.Printf("harvesting dates [%s, %s]\n", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	var g errgroup.Group
	definitionsChan := make(chan []definition_collector.Definition, 100)
	defer close(definitionsChan)
	endSendingDefinitions := make(chan any, 1)
	defer close(endSendingDefinitions)

	go d.StoreDefinitions(ctx, definitionsChan, endSendingDefinitions)

	iterator := 0
	for harvestingDate := startDate; harvestingDate.Before(endDate); harvestingDate = harvestingDate.AddDate(0, 0, iterator) {
		go func(dateCopy time.Time) {
			g.Go(func() error {
				return d.HarvestDefinitionsDate(ctx, dateCopy, definitionsChan)
			})
		}(harvestingDate)
		iterator++
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	<-endSendingDefinitions

	return nil
}

func (d DefinitionController) HarvestDefinitionsDate(ctx context.Context, date time.Time, definitionsChan chan []definition_collector.Definition) error {
	definitions, err := d.DefinitionCollector.GetDefinitions(ctx, date)
	if err != nil {
		return fmt.Errorf("failed to get definitions : %s", err)
	}
	definitionsChan <- definitions
	return nil
}

func (d DefinitionController) StoreDefinitions(ctx context.Context, definitionsChan chan []definition_collector.Definition, endSendingDefeiition chan any) error {
	for {
		select {
		case definitions := <-definitionsChan:
			if err := d.DataHandler.StoreDefinitions(ctx, definitions); err != nil {
				return fmt.Errorf("failed to store definitions : %s", err)
			}
		case <-endSendingDefeiition:
			return nil
		}
	}
}
