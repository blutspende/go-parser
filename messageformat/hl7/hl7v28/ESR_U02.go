package hl7v28

// ESR_U02 - Automated equipment status request
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ESR_U02
type ESR_U02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
}

