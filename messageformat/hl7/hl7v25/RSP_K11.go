package hl7v25

// RSP_K11 - Segment Pattern Response in Response to QBP_Q11
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/RSP_K11
type RSP_K11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

