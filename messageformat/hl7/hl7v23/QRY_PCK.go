package hl7v23

// QRY_PCK - Pathway (goal-oriented) query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/QRY_PCK
type QRY_PCK struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
}

