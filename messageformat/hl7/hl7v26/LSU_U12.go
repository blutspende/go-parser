package hl7v26

// LSU_U12 - Automated Equipment Log/Service Update
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/LSU_U12
type LSU_U12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	EquipmentLogService []EQP `hl7:"TAG=EQP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

