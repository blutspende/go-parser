package hl7v231

// QRY_Q28 - Pharmacy/treatment dispense information query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/QRY_Q28
type QRY_Q28 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

