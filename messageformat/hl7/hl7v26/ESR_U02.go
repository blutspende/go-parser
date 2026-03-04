package hl7v26

// ESR_U02 - Automated Equipment Status Request
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/ESR_U02
type ESR_U02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

