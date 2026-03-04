package hl7v26

// SLN_S34 - Notification of New Sterilization Lot
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SLN_S34
type SLN_S34 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationLotSegment []SLT `hl7:"TAG=SLT"`
}

