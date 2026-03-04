package hl7v26

// QBP_E03_QUERY_INFORMATION - Group struct
type QBP_E03_QUERY_INFORMATION struct {
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
}

// QBP_E03 - HealthCare Services Invoice Status Query
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/QBP_E03
type QBP_E03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential []UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryInformation QBP_E03_QUERY_INFORMATION `hl7:"GROUP"`
}

