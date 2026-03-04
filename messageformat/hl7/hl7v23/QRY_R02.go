package hl7v23

// QRY_R02 - Deferred query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/QRY_R02
type QRY_R02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF"`
}

