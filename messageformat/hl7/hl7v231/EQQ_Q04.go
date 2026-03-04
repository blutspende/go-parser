package hl7v231

// EQQ_Q04 - Embedded query language query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/EQQ_Q04
type EQQ_Q04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EmbeddedQueryLanguageSegment EQL `hl7:"TAG=EQL"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

