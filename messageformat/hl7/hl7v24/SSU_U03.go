package hl7v24

// SSU_U03_SPECIMEN_CONTAINER - Group struct
type SSU_U03_SPECIMEN_CONTAINER struct {
	SpecimenAndContainerDetail SAC `hl7:"TAG=SAC"`
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
}

// SSU_U03 - Specimen status update
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/SSU_U03
type SSU_U03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	SpecimenContainer []SSU_U03_SPECIMEN_CONTAINER `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

