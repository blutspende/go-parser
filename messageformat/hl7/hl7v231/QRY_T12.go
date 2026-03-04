package hl7v231

// QRY_T12 - Document query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/QRY_T12
type QRY_T12 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
}

