package hl7v251

// VXQ_V01 - Vaccination Record Query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/VXQ_V01
type VXQ_V01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

