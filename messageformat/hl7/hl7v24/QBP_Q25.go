package hl7v24

// QBP_Q25 - Query - Personnel information
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QBP_Q25
type QBP_Q25 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

