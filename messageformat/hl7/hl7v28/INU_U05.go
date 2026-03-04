package hl7v28

// INU_U05 - Automated equipment inventory update
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/INU_U05
type INU_U05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	InventoryDetail []INV `hl7:"TAG=INV"`
}

