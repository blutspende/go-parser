package hl7v28

// ESU_U01 - Automated equipment status update
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ESU_U01
type ESU_U01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	InteractionStatusDetail []ISD `hl7:"TAG=ISD;ATR=optional"`
}

