package hl7v26

// QRY_PC9 - Patient Goal Query
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/QRY_PC9
type QRY_PC9 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

