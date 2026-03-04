package hl7v28

// SLR_S28 - Request new sterilization lot
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/SLR_S28
type SLR_S28 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationLot []SLT `hl7:"TAG=SLT"`
}

