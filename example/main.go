package main

import (
	"fmt"
	"strconv"

	"github.com/mineamihai2001/go-cli"
)

func main() {
	manager := cli.Create()

	var totalWorkers int = 20
	var totalRequest int = 1

	manager.Add(&cli.Command{
		Name:        "total",
		Short:       "t",
		Description: "The total number of requests that are going to be fired",
		Required:    true,
		Handler: func(params ...string) {
			total, _ := strconv.Atoi(params[0])
			totalRequest = total
		},
	})

	manager.Add(&cli.Command{
		Name:        "workers",
		Short:       "t",
		Description: "Number of workers which will process the requests. (defaults to 20)",
		Handler: func(params ...string) {
			workers, _ := strconv.Atoi(params[0])
			totalWorkers = workers
		},
	})

	manager.Add(&cli.Command{
		Name:        "url",
		Short:       "u",
		Description: "Url the requests will be made to",
		Handler:     func(params ...string) {},
	})

	manager.Start()

	fmt.Printf("Supplied values: \n\t -totalWorkers=%d \n\t -totalRequests=%d \n", totalWorkers, totalRequest)
}
