package hl7v23

// QRY_T12 - Document query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/QRY_T12
type QRY_T12 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
}

