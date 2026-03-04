package hl7v26

// QRY_R02 - Results of Observation Query
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/QRY_R02
type QRY_R02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF"`
}

