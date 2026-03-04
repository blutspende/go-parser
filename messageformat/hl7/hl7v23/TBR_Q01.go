package hl7v23

// TBR_Q01 - Tabular data response
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/TBR_Q01
type TBR_Q01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK"`
	TableRowDefinition RDF `hl7:"TAG=RDF"`
	TableRowData []RDT `hl7:"TAG=RDT"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

