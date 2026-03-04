package hl7v231

// QRY_PC4 - Problem query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/QRY_PC4
type QRY_PC4 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
}

