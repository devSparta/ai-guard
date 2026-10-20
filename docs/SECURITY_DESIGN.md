# Security Design

## Конвейер безопасности

Независимые detectors можно выполнять конкурентно.

## DLP и Secret Detector

| Компонент | Что ищет |
|---|---|
| DLP Scanner | персональные, финансовые, корпоративные и другие чувствительные данные |
| Secret Detector | API keys, токены, passwords, private keys, connection strings |

Технический secret обычно имеет более высокий severity, поэтому выделен отдельно.

## Источник данных

| Source | Значение |
|---|---|
| `USER_INPUT` | прямой ввод пользователя |
| `RETRIEVED_DOCUMENT` | документ/RAG |
| `TOOL_RESULT` | результат инструмента |
| `LLM_OUTPUT` | ответ модели |
| `EXTERNAL_CONTENT` | внешние данные |

Источник критичен для indirect prompt injection: инструкция из документа не становится доверенной только потому, что документ найден корпоративным поиском.

## Решения

| Decision | Смысл |
|---|---|
| `ALLOW` | разрешить |
| `ALLOW_WITH_MASKING` | разрешить после маскирования |
| `REQUIRE_CONFIRMATION` | потребовать подтверждение |
| `REVIEW` | отправить на ручную проверку |
| `BLOCK` | заблокировать |

## Output Scanner

Финальный ответ LLM также проходит DLP/Secret/System Prompt проверки. Безопасный вход не гарантирует безопасный выход.
