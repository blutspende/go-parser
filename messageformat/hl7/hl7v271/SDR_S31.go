package hl7v271

// SDR_S31 - Request anti-microbial device data
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/SDR_S31
type SDR_S31 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
}

