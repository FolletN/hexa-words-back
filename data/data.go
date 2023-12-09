package data

import (
	"context"
	"cross-words-harverter/interprete"
)

type Data interface {
	StoreSolutions(ctx context.Context, solutions []interprete.Solution) error
}
