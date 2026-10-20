# Контракт студента 4 — LLM / Agent / Tools / DevOps

## LLMProvider

```go
type LLMProvider interface {
    Name() string
    Generate(ctx context.Context, req GenerateRequest) (GenerateResponse, error)
}
```

## Agent Runtime

Agent получает уже проверенный запрос, вызывает LLM, передаёт tool call в Tool Guard и никогда не выполняет tool напрямую. ToolResult обязательно возвращается в Security Pipeline.

## Tool

```go
type Tool interface {
    Definition() ToolDefinition
    Execute(ctx context.Context, input ToolInput) (ToolResult, error)
}
```

Для v1 достаточно `search_documents`, `read_document` и безопасного demo-action `send_demo_email` с confirmation.

CI должен включать gofmt/go vet/go test/go test -race, frontend checks, docker build и OpenAPI validation.
