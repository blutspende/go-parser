package hl7v25

// QRY_PC9 - Patient goal query
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/QRY_PC9
type QRY_PC9 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

