package hl7v231

// VXQ_V01 - Query for vaccination record
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/VXQ_V01
type VXQ_V01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
}

