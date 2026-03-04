package hl7v271

// STI_S30 - Request item
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/STI_S30
type STI_S30 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationLot []SLT `hl7:"TAG=SLT"`
}

