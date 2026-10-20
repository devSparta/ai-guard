package domain

type Evidence struct {
	Kind    string `json:"kind"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Preview string `json:"preview"`
}

type DetectionResult struct {
	Detector    string         `json:"detector"`
	Category    ThreatCategory `json:"category"`
	Detected    bool           `json:"detected"`
	Score       int            `json:"score"`
	Confidence  float64        `json:"confidence"`
	Severity    Severity       `json:"severity"`
	ReasonCodes []string       `json:"reason_codes"`
	Evidence    []Evidence     `json:"evidence"`
}

type SecurityResult struct {
	RiskScore   int               `json:"risk_score"`
	Severity    Severity          `json:"severity"`
	Decision    SecurityDecision  `json:"decision"`
	ReasonCodes []string          `json:"reason_codes"`
	Detections  []DetectionResult `json:"detections,omitempty"`
}
