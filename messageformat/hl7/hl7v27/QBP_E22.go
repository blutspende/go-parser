package hl7v27

// QBP_E22_QUERY - Group struct
type QBP_E22_QUERY struct {
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
}

// QBP_E22 - Query Authorization Request Status
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/QBP_E22
type QBP_E22 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	Query QBP_E22_QUERY `hl7:"GROUP"`
}

