package hl7v24

// EQQ_Q04 - Embedded query language query
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/EQQ_Q04
type EQQ_Q04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EmbeddedQueryLanguage EQL `hl7:"TAG=EQL"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

