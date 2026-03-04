package hl7v27

// SLN_S35 - Notification of sterilization lot deletion
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/SLN_S35
type SLN_S35 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationLot []SLT `hl7:"TAG=SLT"`
}

