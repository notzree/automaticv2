package main

import (
	"context"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

func main() {
	ctx := context.Background()
	g := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)
	resp, err := genkit.Generate(context.Background(), g, ai.WithReturnToolRequests(true))
	if err != nil {
		panic(err)
	}
	for _, toolRequest := range resp.ToolRequests() {
		_ = toolRequest // TODO: Handle tool requests
	}

}
