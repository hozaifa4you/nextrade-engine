package main

import (
	"context"
	"log"
	"os"

	"nextrade-engine/internal/agent"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

func main() {
	chatModel, err := agent.New(context.Background(), agent.Config{
		APIKey:  os.Getenv("OPENROUTER_API_KEY"),
		Model:   envOrDefault("OPENROUTER_MODEL", "openai/gpt-4o-mini"),
		BaseURL: envOrDefault("OPENROUTER_BASE_URL", "https://openrouter.ai/api/v1"),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	events := chatModel.Run(ctx, &adk.TypedAgentInput[*schema.Message]{
		Messages: []*schema.Message{schema.UserMessage("Hello, how are you?")},
	})

	for {
		event, ok := events.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			log.Fatal(event.Err)
		}
		if event.Output != nil && event.Output.MessageOutput != nil {
			message, err := event.Output.MessageOutput.GetMessage()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("AI agent response: ", message.Content)
		}
	}
}

func envOrDefault(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
