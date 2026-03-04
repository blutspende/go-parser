package hl7v26

// SMD_S32_ANTI_MICROBIAL_DEVICE_CYCLE_DATA - Group struct
type SMD_S32_ANTI_MICROBIAL_DEVICE_CYCLE_DATA struct {
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
	AntiMicrobialCycleData []SCD `hl7:"TAG=SCD;ATR=optional"`
}

// SMD_S32 - Request Anti-Microbial Device Cycle Data 
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SMD_S32
type SMD_S32 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	AntiMicrobialDeviceCycleData []SMD_S32_ANTI_MICROBIAL_DEVICE_CYCLE_DATA `hl7:"GROUP"`
}

