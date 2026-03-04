package hl7v28

// LSR_U13 - Automated equipment log/service request
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/LSR_U13
type LSR_U13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	EquipmentLogService []EQP `hl7:"TAG=EQP"`
}

