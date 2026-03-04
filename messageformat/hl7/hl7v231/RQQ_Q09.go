package hl7v231

// RQQ_Q09 - Event replay query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RQQ_Q09
type RQQ_Q09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventReplayQuerySegment ERQ `hl7:"TAG=ERQ"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

