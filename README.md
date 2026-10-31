# Nextrade Engine

An AI agent written in Go.

The agent uses [Eino](https://github.com/cloudwego/eino) with OpenRouter's
OpenAI-compatible API.

## Run

```bash
copy .env.example .env
air
```

For a direct run, export the variables from `.env` first:

```bash
go run ./cmd/agent
```

## Structure

- `cmd/agent` - application entrypoint
- `internal` - private application packages
