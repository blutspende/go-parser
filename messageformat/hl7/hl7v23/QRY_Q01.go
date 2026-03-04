package hl7v23

// QRY_Q01 - Query sent for immediate response
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/QRY_Q01
type QRY_Q01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

