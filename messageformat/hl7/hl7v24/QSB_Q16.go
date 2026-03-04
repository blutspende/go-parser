package hl7v24

// QSB_Q16 - Create subscription
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QSB_Q16
type QSB_Q16 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

