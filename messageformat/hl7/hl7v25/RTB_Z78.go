package hl7v25

// RTB_Z78_ROW_DEFINITION - Group struct
type RTB_Z78_ROW_DEFINITION struct {
	TableRowDefinition RDF `hl7:"TAG=RDF"`
	TableRowData []RDT `hl7:"TAG=RDT;ATR=optional"`
}

// RTB_Z78 - Query By Example - Tabular Patient List Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/RTB_Z78
type RTB_Z78 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	RowDefinition RTB_Z78_ROW_DEFINITION `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

