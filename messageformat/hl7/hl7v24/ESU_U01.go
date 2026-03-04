package hl7v24

// ESU_U01 - Automated equipment status update
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ESU_U01
type ESU_U01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	InteractionStatusDetail []ISD `hl7:"TAG=ISD;ATR=optional"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

