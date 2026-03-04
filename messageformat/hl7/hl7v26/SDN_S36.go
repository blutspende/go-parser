package hl7v26

// SDN_S36_ANTI_MICROBIAL_DEVICE_DATA - Group struct
type SDN_S36_ANTI_MICROBIAL_DEVICE_DATA struct {
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
	AntiMicrobialCycleData []SCD `hl7:"TAG=SCD;ATR=optional"`
}

// SDN_S36 - Notification of Anti-Microbial Device Data
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SDN_S36
type SDN_S36 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	AntiMicrobialDeviceData []SDN_S36_ANTI_MICROBIAL_DEVICE_DATA `hl7:"GROUP"`
}

