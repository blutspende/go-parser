package hl7v24

// LSU_U12 - Automated equipment log/service update
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/LSU_U12
type LSU_U12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	EquipmentLogService []EQP `hl7:"TAG=EQP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

