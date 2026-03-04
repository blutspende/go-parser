package hl7v25

// MFN_M05_MF_LOCATION - Group struct
type MFN_M05_MF_LOCATION struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	LocationIdentification LOC `hl7:"TAG=LOC"`
	LocationCharacteristic []LCH `hl7:"TAG=LCH;ATR=optional"`
	LocationRelationship []LRL `hl7:"TAG=LRL;ATR=optional"`
	MfLocDept []MFN_M05_MF_LOCATION_MF_LOC_DEPT `hl7:"GROUP"`
}

// MFN_M05_MF_LOCATION_MF_LOC_DEPT - Group struct
type MFN_M05_MF_LOCATION_MF_LOC_DEPT struct {
	LocationDepartment LDP `hl7:"TAG=LDP"`
	LocationCharacteristic []LCH `hl7:"TAG=LCH;ATR=optional"`
	LocationChargeCode []LCC `hl7:"TAG=LCC;ATR=optional"`
}

// MFN_M05 - Master files notification - Patient location master file
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/MFN_M05
type MFN_M05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfLocation []MFN_M05_MF_LOCATION `hl7:"GROUP"`
}

