package hl7v26

// INR_U06 - Automated Equipment Inventory Request
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/INR_U06
type INR_U06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	InventoryDetail []INV `hl7:"TAG=INV"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

