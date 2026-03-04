package hl7v27

// EAC_U07_COMMAND - Group struct
type EAC_U07_COMMAND struct {
	EquipmentCommand ECD `hl7:"TAG=ECD"`
	TimingQuantity TQ1 `hl7:"TAG=TQ1;ATR=optional"`
	Specimen_container EAC_U07_COMMAND_SPECIMEN_CONTAINER `hl7:"GROUP;ATR=optional"`
	ClearNotification CNS `hl7:"TAG=CNS;ATR=optional"`
}

// EAC_U07_COMMAND_SPECIMEN_CONTAINER - Group struct
type EAC_U07_COMMAND_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	Specimen []SPM `hl7:"TAG=SPM;ATR=optional"`
}

// EAC_U07 - Automated equipment command
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/EAC_U07
type EAC_U07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	Command []EAC_U07_COMMAND `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

