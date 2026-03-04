package hl7v27

// QBP_Z73 - Information about Phone Calls
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/QBP_Z73
type QBP_Z73 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
}

