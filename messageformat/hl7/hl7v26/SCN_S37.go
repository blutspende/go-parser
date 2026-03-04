package hl7v26

// SCN_S37_ANTI_MICROBIAL_DEVICE_CYCLE_DATA - Group struct
type SCN_S37_ANTI_MICROBIAL_DEVICE_CYCLE_DATA struct {
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
	AntiMicrobialCycleData []SCD `hl7:"TAG=SCD;ATR=optional"`
}

// SCN_S37 - Notification of Anti-Microbial Device Cycle Data
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SCN_S37
type SCN_S37 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	AntiMicrobialDeviceCycleData []SCN_S37_ANTI_MICROBIAL_DEVICE_CYCLE_DATA `hl7:"GROUP"`
}

