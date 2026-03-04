package hl7v24

// SSR_U04 - Specimen status request
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/SSR_U04
type SSR_U04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	SpecimenAndContainerDetail []SAC `hl7:"TAG=SAC"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

