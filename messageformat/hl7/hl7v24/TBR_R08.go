package hl7v24

// TBR_R08 - Tabular data response
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/TBR_R08
type TBR_R08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	TableRowDefinition RDF `hl7:"TAG=RDF"`
	TableRowData []RDT `hl7:"TAG=RDT"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

