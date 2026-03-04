package hl7v24

// QVR_Q17 - Query for previous events
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QVR_Q17
type QVR_Q17 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

