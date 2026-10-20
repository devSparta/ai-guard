# Contract Freeze v1

Фиксируем:

- `/api/v1`;
- JSON `snake_case`;
- UUID string;
- RFC3339 UTC;
- SecurityDecision;
- ThreatCategory;
- SecuritySourceType;
- DetectionResult;
- SecurityResult;
- общий API error;
- interfaces Detector, LLMProvider и Tool.

Можно добавлять optional fields, endpoints, reason codes, detectors и tools.

Breaking change: удаление/переименование поля, смена типа, URL, смысла enum или интерфейса другого участника. Для такого изменения нужны Issue, согласование и обновление документации до реализации.
