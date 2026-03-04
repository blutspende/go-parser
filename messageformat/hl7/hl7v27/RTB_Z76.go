package hl7v27

// RTB_Z76_ROW_DEFINITION - Group struct
type RTB_Z76_ROW_DEFINITION struct {
	TableRowDefinition RDF `hl7:"TAG=RDF"`
	TableRowData []RDT `hl7:"TAG=RDT;ATR=optional"`
}

// RTB_Z76 - Tabular Patient List (Response)
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/RTB_Z76
type RTB_Z76 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Row_definition RTB_Z76_ROW_DEFINITION `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

