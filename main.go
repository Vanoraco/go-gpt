package main

import (
	"context"
	"fmt"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	apiKey := os.Getenv("OPEN_API")

	ctx := context.Background()

	client := gpt3.NewClient(apiKey)

	request := gpt3.CompletionRequest{
		Prompt: []string{"How many coffees should i drink per day?"},
	}

	resp, err := client.Completion(ctx, request)

	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("Answer:\n %s\n", resp.Choices[0].Text)
	}

}
