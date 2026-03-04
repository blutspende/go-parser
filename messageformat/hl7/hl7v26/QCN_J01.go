package hl7v26

// QCN_J01 - Cancel Query
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/QCN_J01
type QCN_J01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryIdentification QID `hl7:"TAG=QID"`
}

