package main

type ParsedQuery struct {
	Type      string    `json:"type"`
	Index     string    `json:"index"`
	Condition Condition `json:"condition,omitempty"`
	OrderBy   []Order   `json:"order_by,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Offset    int       `json:"offset,omitempty"`
}

type Condition struct {
	Operator       string
	ConditionParts []ConditionPart
}

type ConditionPart struct {
	Type      string     `json:"type"`
	Field     string     `json:"field,omitempty"`
	Operator  string     `json:"operator,omitempty"`
	Value     any        `json:"value,omitempty"`
	Condition *Condition `json:"condition,omitempty"`
	Negate    bool       `json:"negate,omitempty"`
}

type Order struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}
