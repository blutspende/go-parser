package hl7v23

// VXQ_V01 - Query for vaccination record
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/VXQ_V01
type VXQ_V01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
}

