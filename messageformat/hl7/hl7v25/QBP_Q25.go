package hl7v25

// QBP_Q25 - Query - Personnel information
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/QBP_Q25
type QBP_Q25 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

