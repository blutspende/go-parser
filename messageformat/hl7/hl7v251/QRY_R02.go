package hl7v251

// QRY_R02 - Query for results of observation
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QRY_R02
type QRY_R02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF"`
}

