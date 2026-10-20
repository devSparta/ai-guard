# AI Guard

**AI Guard** — учебный программный продукт для защиты корпоративных LLM-приложений и AI-агентов.

Система находится между пользователем, языковой моделью и инструментами агента. Её задача — проверять входные данные, внешние документы, вызовы инструментов и ответы модели, чтобы снизить риск prompt injection, утечек конфиденциальных данных и выполнения опасных действий.

Проект разрабатывается командой из четырёх студентов направления **«Программная инженерия»**.

## Основной стек

| Область | Технология |
|---|---|
| Backend | Go |
| HTTP API | `net/http` на старте, при необходимости Gin |
| База данных | PostgreSQL |
| Frontend | React + TypeScript |
| API-контракт | OpenAPI 3.0 |
| Контейнеризация | Docker / Docker Compose |
| CI | GitHub Actions |
| Опциональный ML | отдельный Python-сервис |

## Кто за что отвечает

| Участник | Основная зона |
|---|---|
| **Студент 1** | Go backend, архитектура, БД, API, Auth/RBAC, orchestration |
| **Студент 2** | Security Engine: Prompt Injection, DLP, Secret Detector, Risk Engine |
| **Студент 3** | React frontend, Dashboard, Chat, Security Events |
| **Студент 4** | LLM Provider, AI Agent, Tools, Tool Guard integration, Docker, CI |

Подробное распределение: [`docs/TEAM_AND_OWNERSHIP.md`](docs/TEAM_AND_OWNERSHIP.md).

## Документация

- [`docs/PROJECT_OVERVIEW.md`](docs/PROJECT_OVERVIEW.md) — что строим и где границы продукта;
- [`docs/SECURITY_DESIGN.md`](docs/SECURITY_DESIGN.md) — защитный конвейер;
- [`docs/DECOMPOSITION.md`](docs/DECOMPOSITION.md) — декомпозиция;
- [`docs/TEAM_AND_OWNERSHIP.md`](docs/TEAM_AND_OWNERSHIP.md) — зоны ответственности;
- [`docs/contracts/`](docs/contracts/) — зафиксированные контракты;
- [`docs/spec/openapi.yaml`](docs/spec/openapi.yaml) — HTTP API;
