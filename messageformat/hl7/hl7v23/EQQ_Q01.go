package hl7v23

// EQQ_Q01 - Embedded query language query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/EQQ_Q01
type EQQ_Q01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EmbeddedQueryLanguage EQL `hl7:"TAG=EQL"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

