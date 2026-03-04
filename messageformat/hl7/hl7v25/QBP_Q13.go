package hl7v25

// QBP_Q13 - Query by Parameter Requesting an RTB Tabular Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/QBP_Q13
type QBP_Q13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

