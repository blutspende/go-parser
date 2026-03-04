package hl7v231

// MFQ_M04 - Query for Master File Record - Charge description 
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFQ_M04
type MFQ_M04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

