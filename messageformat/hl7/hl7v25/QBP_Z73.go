package hl7v25

// QBP_Z73 - Information about Phone Calls Query
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/QBP_Z73
type QBP_Z73 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
}

