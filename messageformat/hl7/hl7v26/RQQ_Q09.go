package hl7v26

// RQQ_Q09 - Event Replay Query
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/RQQ_Q09
type RQQ_Q09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventReplayQuery ERQ `hl7:"TAG=ERQ"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

