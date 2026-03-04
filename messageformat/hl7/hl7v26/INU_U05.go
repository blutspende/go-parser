package hl7v26

// INU_U05 - Automated Equipment Inventory Update
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/INU_U05
type INU_U05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	InventoryDetail []INV `hl7:"TAG=INV"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

