package hl7v231

// QRY_A19 - Patient query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/QRY_A19
type QRY_A19 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
}

