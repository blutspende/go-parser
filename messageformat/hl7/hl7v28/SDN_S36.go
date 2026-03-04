package hl7v28

// SDN_S36_ANTI_MICROBIAL_DEVICE_DATA - Group struct
type SDN_S36_ANTI_MICROBIAL_DEVICE_DATA struct {
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
	AntiMicrobialCycleData []SCD `hl7:"TAG=SCD;ATR=optional"`
}

// SDN_S36 - Notification of anti-microbial device data
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/SDN_S36
type SDN_S36 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	AntiMicrobial_device_data SDN_S36_ANTI_MICROBIAL_DEVICE_DATA `hl7:"GROUP"`
}

