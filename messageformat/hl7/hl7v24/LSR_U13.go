package hl7v24

// LSR_U13 - Automated equipment log/service request
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/LSR_U13
type LSR_U13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	EquipmentLogService []EQP `hl7:"TAG=EQP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

