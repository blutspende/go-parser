package hl7v251

// QBP_Z99 - Query by Parameter - Who Am I
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QBP_Z99
type QBP_Z99 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

