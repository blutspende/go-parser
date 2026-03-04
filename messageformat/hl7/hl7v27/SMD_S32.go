package hl7v27

// SMD_S32 - Request anti-microbial device cycle data
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/SMD_S32
type SMD_S32 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
}

