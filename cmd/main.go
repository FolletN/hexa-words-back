package main

import (
	"context"
	"cross-words/internal/view/api"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Initializing server")
	server := api.NewClient(api.ServerConfiguration{
		Port: 8080,
	})

	if err := server.Serve(ctx); err != nil {
		panic(err)
	}
}
