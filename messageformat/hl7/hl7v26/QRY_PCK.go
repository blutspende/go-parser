package hl7v26

// QRY_PCK - Patient Pathway (Goal-Oriented) Query
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/QRY_PCK
type QRY_PCK struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

