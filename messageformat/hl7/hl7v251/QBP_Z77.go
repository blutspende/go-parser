package hl7v251

// QBP_Z77 - Query By Example - Tabular Patient List Query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QBP_Z77
type QBP_Z77 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	PatientIdentification PID `hl7:"TAG=PID"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

