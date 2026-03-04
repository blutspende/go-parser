package hl7v26

// EAR_U08_COMMAND_RESPONSE - Group struct
type EAR_U08_COMMAND_RESPONSE struct {
	EquipmentCommand ECD `hl7:"TAG=ECD"`
	SpecimenContainer EAR_U08_COMMAND_RESPONSE_SPECIMEN_CONTAINER `hl7:"GROUP;ATR=optional"`
	EquipmentCommandResponse ECR `hl7:"TAG=ECR"`
}

// EAR_U08_COMMAND_RESPONSE_SPECIMEN_CONTAINER - Group struct
type EAR_U08_COMMAND_RESPONSE_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	Specimen []SPM `hl7:"TAG=SPM;ATR=optional"`
}

// EAR_U08 - Automated Equipment Response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/EAR_U08
type EAR_U08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	CommandResponse []EAR_U08_COMMAND_RESPONSE `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

