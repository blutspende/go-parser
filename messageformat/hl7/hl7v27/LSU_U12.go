package hl7v27

// LSU_U12 - Automated equipment log/service update
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/LSU_U12
type LSU_U12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	EquipmentLogService []EQP `hl7:"TAG=EQP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

