package hl7v24

// QRY_R02 - Query for results of observation
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QRY_R02
type QRY_R02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF"`
}

