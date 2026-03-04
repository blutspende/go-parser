package hl7v26

// SDR_S31_ANTI_MICROBIAL_DEVICE_DATA - Group struct
type SDR_S31_ANTI_MICROBIAL_DEVICE_DATA struct {
	SterilizationDeviceData SDD `hl7:"TAG=SDD"`
	AntiMicrobialCycleData []SCD `hl7:"TAG=SCD;ATR=optional"`
}

// SDR_S31 - Request Sterilization Anti-Microbial Device Data
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SDR_S31
type SDR_S31 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	AntiMicrobialDeviceData []SDR_S31_ANTI_MICROBIAL_DEVICE_DATA `hl7:"GROUP"`
}

