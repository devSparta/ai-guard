package domain

type SecurityDecision string

const (
	SecurityDecisionAllow               SecurityDecision = "ALLOW"
	SecurityDecisionAllowWithMasking    SecurityDecision = "ALLOW_WITH_MASKING"
	SecurityDecisionRequireConfirmation SecurityDecision = "REQUIRE_CONFIRMATION"
	SecurityDecisionReview              SecurityDecision = "REVIEW"
	SecurityDecisionBlock               SecurityDecision = "BLOCK"
)

type Severity string

const (
	SeverityLow      Severity = "LOW"
	SeverityMedium   Severity = "MEDIUM"
	SeverityHigh     Severity = "HIGH"
	SeverityCritical Severity = "CRITICAL"
)

type SecuritySourceType string

const (
	SourceUserInput         SecuritySourceType = "USER_INPUT"
	SourceRetrievedDocument SecuritySourceType = "RETRIEVED_DOCUMENT"
	SourceToolResult        SecuritySourceType = "TOOL_RESULT"
	SourceLLMOutput         SecuritySourceType = "LLM_OUTPUT"
	SourceExternalContent   SecuritySourceType = "EXTERNAL_CONTENT"
)

type ThreatCategory string

const (
	ThreatPromptInjection          ThreatCategory = "PROMPT_INJECTION"
	ThreatIndirectPromptInjection  ThreatCategory = "INDIRECT_PROMPT_INJECTION"
	ThreatJailbreak                ThreatCategory = "JAILBREAK"
	ThreatSystemPromptExfiltration ThreatCategory = "SYSTEM_PROMPT_EXFILTRATION"
	ThreatPIILeak                  ThreatCategory = "PII_LEAK"
	ThreatSensitiveDataLeak        ThreatCategory = "SENSITIVE_DATA_LEAK"
	ThreatSecretLeak               ThreatCategory = "SECRET_LEAK"
	ThreatPolicyViolation          ThreatCategory = "POLICY_VIOLATION"
	ThreatUnsafeToolCall           ThreatCategory = "UNSAFE_TOOL_CALL"
	ThreatExcessivePrivilege       ThreatCategory = "EXCESSIVE_PRIVILEGE"
	ThreatUntrustedContent         ThreatCategory = "UNTRUSTED_CONTENT"
)
