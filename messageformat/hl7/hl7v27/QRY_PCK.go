package hl7v27

// QRY_PCK - PC/ pathway (goal-oriented) query
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/QRY_PCK
type QRY_PCK struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
}

