package hl7v23

// QRY_Q27 - Pharmacy/treatment administration information query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/QRY_Q27
type QRY_Q27 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

