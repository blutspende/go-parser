package hl7v24

// EAC_U07 - Automated equipment command
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/EAC_U07
type EAC_U07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	EquipmentCommand []ECD `hl7:"TAG=ECD"`
	SpecimenAndContainerDetail SAC `hl7:"TAG=SAC;ATR=optional"`
	ClearNotification CNS `hl7:"TAG=CNS;ATR=optional"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

