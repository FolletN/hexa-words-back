package main

import (
	"context"
	"cross-words-harverter/data/db"
	"cross-words-harverter/handler"
)

func main() {
	ctx := context.Background()
	newHandler := handler.Handler{
		DataHandler: db.NewDB(),
	}
	err := newHandler.Handle(ctx)
	if err != nil {
		panic(err)
	}
}
