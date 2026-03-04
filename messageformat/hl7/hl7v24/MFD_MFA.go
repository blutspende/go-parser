package hl7v24

// MFD_MFA - Master files delayed application acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFD_MFA
type MFD_MFA struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

