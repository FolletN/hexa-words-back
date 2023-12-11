package model

import (
	"context"

	"cross-words/internal/controller/definition_collector"
)

type Data interface {
	StoreDefinitions(ctx context.Context, solutions []definition_collector.Definition) error
}
