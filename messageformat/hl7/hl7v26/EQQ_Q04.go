package hl7v26

// EQQ_Q04 - Embedded Query Language Query
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/EQQ_Q04
type EQQ_Q04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EmbeddedQueryLanguage EQL `hl7:"TAG=EQL"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

