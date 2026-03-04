package hl7v251

// QRY_PCK - Patient pathway (goal-oriented) query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QRY_PCK
type QRY_PCK struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

