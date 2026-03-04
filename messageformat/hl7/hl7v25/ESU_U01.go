package hl7v25

// ESU_U01 - Automated equipment status update
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ESU_U01
type ESU_U01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	InteractionStatusDetail []ISD `hl7:"TAG=ISD;ATR=optional"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

