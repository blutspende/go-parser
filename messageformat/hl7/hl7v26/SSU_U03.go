package hl7v26

// SSU_U03_SPECIMEN_CONTAINER - Group struct
type SSU_U03_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	Specimen []SSU_U03_SPECIMEN_CONTAINER_SPECIMEN `hl7:"GROUP;ATR=optional"`
}

// SSU_U03_SPECIMEN_CONTAINER_SPECIMEN - Group struct
type SSU_U03_SPECIMEN_CONTAINER_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// SSU_U03 - Specimen Status Update
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SSU_U03
type SSU_U03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	SpecimenContainer []SSU_U03_SPECIMEN_CONTAINER `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

