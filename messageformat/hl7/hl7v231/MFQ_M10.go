package hl7v231

// MFQ_M10 - Query for Master File Record - Test /observation batteries
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFQ_M10
type MFQ_M10 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

