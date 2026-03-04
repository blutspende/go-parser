package hl7v231

// MFN_M05_MF_LOCATION - Group struct
type MFN_M05_MF_LOCATION struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
	LocationIdentificationSegment LOC `hl7:"TAG=LOC"`
	LocationCharacteristicSegment []LCH `hl7:"TAG=LCH;ATR=optional"`
	LocationRelationshipSegment []LRL `hl7:"TAG=LRL;ATR=optional"`
	MfLocDept []MFN_M05_MF_LOCATION_MF_LOC_DEPT `hl7:"GROUP"`
}

// MFN_M05_MF_LOCATION_MF_LOC_DEPT - Group struct
type MFN_M05_MF_LOCATION_MF_LOC_DEPT struct {
	LocationDepartmentSegment LDP `hl7:"TAG=LDP"`
	LocationCharacteristicSegment []LCH `hl7:"TAG=LCH;ATR=optional"`
	LocationChargeCodeSegment []LCC `hl7:"TAG=LCC;ATR=optional"`
}

// MFN_M05 - Patient location master file
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFN_M05
type MFN_M05 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfLocation []MFN_M05_MF_LOCATION `hl7:"GROUP"`
}

