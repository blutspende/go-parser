package hl7v26

// QBP_Q23 - Get Corresponding Identifiers Query
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/QBP_Q23
type QBP_Q23 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

