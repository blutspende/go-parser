package hl7v27

// SSR_U04_SPECIMEN_CONTAINER - Group struct
type SSR_U04_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	Specimen []SPM `hl7:"TAG=SPM;ATR=optional"`
}

// SSR_U04 - Specimen status request
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/SSR_U04
type SSR_U04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	Specimen_container []SSR_U04_SPECIMEN_CONTAINER `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

