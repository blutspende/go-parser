package hl7v26

// STC_S33 - Notification of Sterilization Configuration
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/STC_S33
type STC_S33 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	SterilizerConfiguration []SCP `hl7:"TAG=SCP"`
}

