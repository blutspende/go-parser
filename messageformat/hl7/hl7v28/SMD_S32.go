package hl7v28

// SMD_S32_ANTI_MICROBIAL_DEVICE_CYCLE_DATA - Group struct
type SMD_S32_ANTI_MICROBIAL_DEVICE_CYCLE_DATA struct {
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
	AntiMicrobialCycleData []SCD `hl7:"TAG=SCD;ATR=optional"`
}

// SMD_S32 - Request anti-microbial device cycle data
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/SMD_S32
type SMD_S32 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	AntiMicrobial_device_cycle_data SMD_S32_ANTI_MICROBIAL_DEVICE_CYCLE_DATA `hl7:"GROUP"`
}

