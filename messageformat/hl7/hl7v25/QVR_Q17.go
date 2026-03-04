package hl7v25

// QVR_Q17 - Query for Previous Events
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/QVR_Q17
type QVR_Q17 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

