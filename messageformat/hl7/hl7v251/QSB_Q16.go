package hl7v251

// QSB_Q16 - Create Subscription
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QSB_Q16
type QSB_Q16 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

