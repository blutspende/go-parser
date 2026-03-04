package hl7v271

// QSX_J02 - Cancel subscription/acknowledge message
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/QSX_J02
type QSX_J02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryIdentification QID `hl7:"TAG=QID"`
}

