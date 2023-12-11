package collector

import (
	"context"
	"hexacrosswords/internal/db"
	"time"
)

type Collector interface {
	GetDefinitions(ctx context.Context, date time.Time) ([]db.Definition, error)
}
