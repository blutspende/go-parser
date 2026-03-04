package hl7v251

// QBP_Q31 - Dispense History Query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QBP_Q31
type QBP_Q31 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

