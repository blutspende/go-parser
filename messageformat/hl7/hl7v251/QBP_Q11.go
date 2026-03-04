package hl7v251

// QBP_Q11 - Query by Parameter Requesting an RSP Segment Pattern Response
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QBP_Q11
type QBP_Q11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

