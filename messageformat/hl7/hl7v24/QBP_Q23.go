package hl7v24

// QBP_Q23 - Query - Get Corresponding Identifiers
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QBP_Q23
type QBP_Q23 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

