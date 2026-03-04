package hl7v24

// QBP_Q15 - Query by parameter/display response
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QBP_Q15
type QBP_Q15 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

