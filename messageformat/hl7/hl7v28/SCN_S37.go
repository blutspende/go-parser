package hl7v28

// SCN_S37_ANTI_MICROBIAL_DEVICE_CYCLE_DATA - Group struct
type SCN_S37_ANTI_MICROBIAL_DEVICE_CYCLE_DATA struct {
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
	AntiMicrobialCycleData []SCD `hl7:"TAG=SCD;ATR=optional"`
}

// SCN_S37 - Notification of anti-microbial device cycle data
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/SCN_S37
type SCN_S37 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	AntiMicrobial_device_cycle_data SCN_S37_ANTI_MICROBIAL_DEVICE_CYCLE_DATA `hl7:"GROUP"`
}

