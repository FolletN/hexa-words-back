package main

import (
	"context"
	"cross-words-harverter/db"
	"cross-words-harverter/httpclient"
	"cross-words-harverter/interprete"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	date := time.Now()
	for date.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)) {
		date = date.AddDate(0, 0, -1)
		stringDate := date.Format("020106")
		fmt.Printf("Processing date %v\n", stringDate)
		time.Sleep(time.Second / 100)
		data, err := httpclient.GetData(ctx, stringDate)
		if err != nil {
			fmt.Printf("Error while getting data : %v\n", err.Error())
			panic(nil)
		}

		solutions := interprete.NewInterpretor(data).Interprete()
		if len(solutions) == 0 {
			fmt.Printf("ERROR\n\n\n")
		}

		for _, solution := range solutions {
			if err := db.InsertWordDefinition(solution); err != nil {
				fmt.Printf("Error while storing data : %v\n", err.Error())
			}
		}
	}

}
