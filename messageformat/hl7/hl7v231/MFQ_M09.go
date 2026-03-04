package hl7v231

// MFQ_M09 - Query for Master File Record - Test/Observation (Categorical)
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFQ_M09
type MFQ_M09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

