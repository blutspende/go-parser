package hl7v231

// QRY_R02 - Query for results of observation
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/QRY_R02
type QRY_R02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF"`
}

