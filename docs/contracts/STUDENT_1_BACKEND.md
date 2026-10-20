# Контракт студента 1 — Backend

Студент 1 отвечает за интеграционный каркас проекта.

Backend гарантирует остальным стабильный `/api/v1`, единые domain types, Auth/RBAC, request/trace IDs, persistence, orchestration, Security Events и OpenAPI.

Backend ожидает от Security:

```go
type Detector interface {
    Name() string
    Analyze(ctx context.Context, input SecurityInput) (DetectionResult, error)
}
```

Backend ожидает от LLM:

```go
type LLMProvider interface {
    Name() string
    Generate(ctx context.Context, req GenerateRequest) (GenerateResponse, error)
}
```

Backend ожидает от Tool:

```go
type Tool interface {
    Definition() ToolDefinition
    Execute(ctx context.Context, input ToolInput) (ToolResult, error)
}
```

Backend не должен знать внутренний алгоритм detector, зависеть от конкретного LLM vendor или разрешать agent выполнять tools в обход Tool Guard.
