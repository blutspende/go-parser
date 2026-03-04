package hl7v26

// MFN_M15_MF_INV_ITEM - Group struct
type MFN_M15_MF_INV_ITEM struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	InventoryItemMaster IIM `hl7:"TAG=IIM"`
}

// MFN_M15 - Inventory Item Master File Notification
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/MFN_M15
type MFN_M15 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfInvItem []MFN_M15_MF_INV_ITEM `hl7:"GROUP"`
}

