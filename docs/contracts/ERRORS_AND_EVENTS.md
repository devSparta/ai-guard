# Ошибки и события

## API error

```json
{
  "error": {
    "code": "SECURITY_BLOCKED",
    "message": "request blocked by security policy",
    "request_id": "uuid",
    "details": {}
  }
}
```

Базовые codes: `VALIDATION_ERROR`, `UNAUTHORIZED`, `FORBIDDEN`, `NOT_FOUND`, `CONFLICT`, `RATE_LIMITED`, `SECURITY_BLOCKED`, `APPROVAL_REQUIRED`, `APPROVAL_EXPIRED`, `LLM_UNAVAILABLE`, `TOOL_NOT_FOUND`, `TOOL_DENIED`, `TOOL_TIMEOUT`, `DATABASE_ERROR`, `INTERNAL_ERROR`.

Security Event отвечает на вопрос «какая угроза обнаружена?», Audit Event — «кто и что сделал в системе?».
