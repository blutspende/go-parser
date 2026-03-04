package hl7v23

// RQQ_Q01 - Event replay query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RQQ_Q01
type RQQ_Q01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventReplayQuerySegment ERQ `hl7:"TAG=ERQ"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

