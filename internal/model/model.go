package model

import (
	"context"
	"cross-words/internal/controler/interprete"
)

type Data interface {
	StoreSolutions(ctx context.Context, solutions []interprete.Solution) error
}
