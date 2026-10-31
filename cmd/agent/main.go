package main

import (
	"context"
	"log"
	"os"

	"nextrade-engine/internal/agent"
)

func main() {
	_, err := agent.New(context.Background(), agent.Config{
		APIKey:  os.Getenv("OPENROUTER_API_KEY"),
		Model:   envOrDefault("OPENROUTER_MODEL", "openai/gpt-4o-mini"),
		BaseURL: envOrDefault("OPENROUTER_BASE_URL", "https://openrouter.ai/api/v1"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("AI agent initialized")
}

func envOrDefault(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
