# Общие типы

## SecurityDecision

`ALLOW`, `ALLOW_WITH_MASKING`, `REQUIRE_CONFIRMATION`, `REVIEW`, `BLOCK`.

## Severity

`LOW`, `MEDIUM`, `HIGH`, `CRITICAL`.

## SecuritySourceType

`USER_INPUT`, `RETRIEVED_DOCUMENT`, `TOOL_RESULT`, `LLM_OUTPUT`, `EXTERNAL_CONTENT`.

## ThreatCategory

`PROMPT_INJECTION`, `INDIRECT_PROMPT_INJECTION`, `JAILBREAK`, `SYSTEM_PROMPT_EXFILTRATION`, `PII_LEAK`, `SENSITIVE_DATA_LEAK`, `SECRET_LEAK`, `POLICY_VIOLATION`, `UNSAFE_TOOL_CALL`, `EXCESSIVE_PRIVILEGE`, `UNTRUSTED_CONTENT`.

## DetectionResult

```json
{
  "detector": "prompt_injection",
  "category": "PROMPT_INJECTION",
  "detected": true,
  "score": 91,
  "confidence": 0.96,
  "severity": "HIGH",
  "reason_codes": ["PI_INSTRUCTION_OVERRIDE"],
  "evidence": []
}
```

## SecurityResult

```json
{
  "risk_score": 91,
  "severity": "CRITICAL",
  "decision": "BLOCK",
  "reason_codes": ["PI_INSTRUCTION_OVERRIDE"],
  "detections": []
}
```
