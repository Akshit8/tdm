package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Akshit8/tdm/pkg/openapi3"
)

func main() {
	client, err := openapi3.NewClientWithResponses("http://0.0.0.0:8000")
	if err != nil {
		log.Fatalln("Couldn't create client: %w", err)
	}

	newPrtStr := func(s string) *string {
		return &s
	}

	newPtrTime := func(t time.Time) *time.Time {
		return &t
	}

	priority := openapi3.Priority_high

	resp, err := client.CreateTaskWithResponse(context.Background(),
		openapi3.CreateTaskJSONRequestBody{
			Dates: &openapi3.Dates{
				Start: newPtrTime(time.Now()),
				Due:   newPtrTime(time.Now().Add(time.Hour)),
			},
			Description: newPrtStr("complete this microservice"),
			Priority:    &priority,
		},
	)
	if err != nil {
		log.Fatalln("Couldn't create task %w", err)
	}

	fmt.Printf("New Task\n\tID: %s\n", *resp.JSON201.Task.Id)
	fmt.Printf("\tDescription: %s\n", *resp.JSON201.Task.Description)
	fmt.Printf("\tPriority: %s\n", *resp.JSON201.Task.Priority)
	fmt.Printf("\tStart: %s\n", *resp.JSON201.Task.Dates.Start)
	fmt.Printf("\tDue: %s\n", *resp.JSON201.Task.Dates.Due)
	fmt.Printf("\tIsDone: %t\n", *resp.JSON201.Task.IsDone)
}
