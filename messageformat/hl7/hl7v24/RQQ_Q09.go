package hl7v24

// RQQ_Q09 - Event replay query
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RQQ_Q09
type RQQ_Q09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventReplayQuery ERQ `hl7:"TAG=ERQ"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

