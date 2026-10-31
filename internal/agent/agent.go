package agent

import (
	"context"
	"errors"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/adk"
)

type Config struct {
	APIKey  string
	Model   string
	BaseURL string
}

func New(ctx context.Context, cfg Config) (*adk.ChatModelAgent, error) {
	if cfg.APIKey == "" {
		return nil, errors.New("OPENROUTER_API_KEY is required")
	}

	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  cfg.APIKey,
		Model:   cfg.Model,
		BaseURL: cfg.BaseURL,
	})
	if err != nil {
		return nil, err
	}

	return adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Model: chatModel,
	})
}
