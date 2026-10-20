# Декомпозиция продукта

## Техническая декомпозиция

| Блок | Получает | Возвращает | Не должен зависеть от |
|---|---|---|---|
| HTTP API | HTTP request | HTTP response | конкретного LLM |
| Security Pipeline | `SecurityInput` | `SecurityResult` | React |
| Detector | текст + metadata | `DetectionResult` | PostgreSQL |
| Risk Engine | detections | score/decision | transport layer |
| LLM Provider | prompt/messages/tools | model response | Security DB |
| Tool Guard | tool call + context | authorization | frontend |
| Tool | validated input | `ToolResult` | UI |
| Event Service | security event | persisted event | LLM implementation |

Контрактная граница нужна для того, чтобы каждый участник мог реализовывать свою часть независимо.
