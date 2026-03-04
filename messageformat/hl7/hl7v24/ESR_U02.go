package hl7v24

// ESR_U02 - Automated equipment status request
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ESR_U02
type ESR_U02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

