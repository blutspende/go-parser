package hl7v271

// STC_S33 - Notification of sterilization configuration
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/STC_S33
type STC_S33 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizerConfigurationAntiMicrobialDevices []SCP `hl7:"TAG=SCP"`
}

