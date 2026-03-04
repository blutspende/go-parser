package hl7v26

// STI_S30 - Request Item
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/STI_S30
type STI_S30 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationLotSegment []SLT `hl7:"TAG=SLT"`
}

