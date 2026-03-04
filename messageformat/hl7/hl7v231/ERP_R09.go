package hl7v231

// ERP_R09 - Event replay response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ERP_R09
type ERP_R09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK"`
	EventReplayQuerySegment ERQ `hl7:"TAG=ERQ"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

