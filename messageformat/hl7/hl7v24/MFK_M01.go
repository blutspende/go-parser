package hl7v24

// MFK_M01 - Master file application acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFK_M01
type MFK_M01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

