package hl7v26

// QSX_J02 - Cancel Subscription
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/QSX_J02
type QSX_J02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryIdentification QID `hl7:"TAG=QID"`
}

