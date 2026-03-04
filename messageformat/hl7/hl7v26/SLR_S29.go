package hl7v26

// SLR_S29 - Request Sterilization Lot Deletion
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SLR_S29
type SLR_S29 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationLotSegment []SLT `hl7:"TAG=SLT"`
}

