package hl7v231

// QRY_P04 - Generate bills and accounts receivable statements 
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/QRY_P04
type QRY_P04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

