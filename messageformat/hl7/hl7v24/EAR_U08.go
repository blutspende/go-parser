package hl7v24

// EAR_U08_COMMAND_RESPONSE - Group struct
type EAR_U08_COMMAND_RESPONSE struct {
	EquipmentCommand ECD `hl7:"TAG=ECD"`
	SpecimenAndContainerDetail SAC `hl7:"TAG=SAC;ATR=optional"`
	EquipmentCommandResponse ECR `hl7:"TAG=ECR"`
}

// EAR_U08 - Automated equipment response
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/EAR_U08
type EAR_U08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	CommandResponse []EAR_U08_COMMAND_RESPONSE `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

