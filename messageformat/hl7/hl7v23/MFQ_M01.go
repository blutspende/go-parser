package hl7v23

// MFQ_M01 - Query for Master File Record - Not otherwise specified 
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MFQ_M01
type MFQ_M01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

