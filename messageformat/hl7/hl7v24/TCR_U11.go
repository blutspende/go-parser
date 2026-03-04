package hl7v24

// TCR_U11 - Automated equipment test code settings request
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/TCR_U11
type TCR_U11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	TestCodeConfiguration []TCC `hl7:"TAG=TCC"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

