package hl7v26

// SLR_S28 - Request New Sterilization Lot
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SLR_S28
type SLR_S28 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationLotSegment []SLT `hl7:"TAG=SLT"`
}

