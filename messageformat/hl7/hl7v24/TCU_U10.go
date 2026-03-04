package hl7v24

// TCU_U10 - Automated equipment test code settings update
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/TCU_U10
type TCU_U10 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	TestCodeConfiguration []TCC `hl7:"TAG=TCC"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

