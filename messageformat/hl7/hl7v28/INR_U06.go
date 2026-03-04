package hl7v28

// INR_U06 - Automated equipment inventory request
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/INR_U06
type INR_U06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	InventoryDetail []INV `hl7:"TAG=INV"`
}

