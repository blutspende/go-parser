package hl7v26

// ERP_R09 - Event Replay Response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/ERP_R09
type ERP_R09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	EventReplayQuery ERQ `hl7:"TAG=ERQ"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

