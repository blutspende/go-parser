package hl7v231

// TBR_R08 - Tabular data response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/TBR_R08
type TBR_R08 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK"`
	TableRowDefinitionSegment RDF `hl7:"TAG=RDF"`
	TableRowDataSegment []RDT `hl7:"TAG=RDT"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

