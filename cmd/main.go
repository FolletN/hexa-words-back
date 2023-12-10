package main

import (
	"context"
	"cross-words/internal/db"
	"cross-words/internal/helper/handler"
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
