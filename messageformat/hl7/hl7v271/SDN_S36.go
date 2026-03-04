package hl7v271

// SDN_S36 - Notification of anti-microbial device data
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/SDN_S36
type SDN_S36 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
}

