package hl7v27

// MFN_M16_MATERIAL_ITEM_RECORD - Group struct
type MFN_M16_MATERIAL_ITEM_RECORD struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	MaterialItem ITM `hl7:"TAG=ITM"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Sterilization []MFN_M16_MATERIAL_ITEM_RECORD_STERILIZATION `hl7:"GROUP;ATR=optional"`
	Purchasing_vendor []MFN_M16_MATERIAL_ITEM_RECORD_PURCHASING_VENDOR `hl7:"GROUP;ATR=optional"`
	Material_location []MFN_M16_MATERIAL_ITEM_RECORD_MATERIAL_LOCATION `hl7:"GROUP;ATR=optional"`
}

// MFN_M16_MATERIAL_ITEM_RECORD_STERILIZATION - Group struct
type MFN_M16_MATERIAL_ITEM_RECORD_STERILIZATION struct {
	SterilizationParameter STZ `hl7:"TAG=STZ"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// MFN_M16_MATERIAL_ITEM_RECORD_PURCHASING_VENDOR - Group struct
type MFN_M16_MATERIAL_ITEM_RECORD_PURCHASING_VENDOR struct {
	PurchasingVendor VND `hl7:"TAG=VND"`
	Packaging []MFN_M16_MATERIAL_ITEM_RECORD_PURCHASING_VENDOR_PACKAGING `hl7:"GROUP;ATR=optional"`
}

// MFN_M16_MATERIAL_ITEM_RECORD_PURCHASING_VENDOR_PACKAGING - Group struct
type MFN_M16_MATERIAL_ITEM_RECORD_PURCHASING_VENDOR_PACKAGING struct {
	ItemPackaging PKG `hl7:"TAG=PKG"`
	PatientChargeCostCenterExceptions []PCE `hl7:"TAG=PCE;ATR=optional"`
}

// MFN_M16_MATERIAL_ITEM_RECORD_MATERIAL_LOCATION - Group struct
type MFN_M16_MATERIAL_ITEM_RECORD_MATERIAL_LOCATION struct {
	MaterialLocation IVT `hl7:"TAG=IVT"`
	MaterialLot []ILT `hl7:"TAG=ILT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// MFN_M16 - Master File Notification Inventory Item Enhanced
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/MFN_M16
type MFN_M16 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Material_item_record []MFN_M16_MATERIAL_ITEM_RECORD `hl7:"GROUP"`
}

