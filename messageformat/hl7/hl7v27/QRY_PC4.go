package hl7v27

// QRY_PC4 - PC/ problem query
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/QRY_PC4
type QRY_PC4 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
}

