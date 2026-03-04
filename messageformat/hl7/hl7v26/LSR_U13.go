package hl7v26

// LSR_U13 - Automated Equipment Log/Service Request
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/LSR_U13
type LSR_U13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	EquipmentLogService []EQP `hl7:"TAG=EQP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

