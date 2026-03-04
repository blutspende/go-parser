package hl7v24

// MFK_M05_MF_LOCATION - Group struct
type MFK_M05_MF_LOCATION struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	LocationIdentification LOC `hl7:"TAG=LOC"`
	LocationCharacteristic []LCH `hl7:"TAG=LCH;ATR=optional"`
	LocationRelationship []LRL `hl7:"TAG=LRL;ATR=optional"`
	MfLocDept []MFK_M05_MF_LOCATION_MF_LOC_DEPT `hl7:"GROUP"`
}

// MFK_M05_MF_LOCATION_MF_LOC_DEPT - Group struct
type MFK_M05_MF_LOCATION_MF_LOC_DEPT struct {
	LocationDepartment LDP `hl7:"TAG=LDP"`
	LocationCharacteristic []LCH `hl7:"TAG=LCH;ATR=optional"`
	LocationChargeCode []LCC `hl7:"TAG=LCC;ATR=optional"`
}

// MFK_M05 - Master file acknowledgment - Patient location master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFK_M05
type MFK_M05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
	MfLocation []MFK_M05_MF_LOCATION `hl7:"GROUP"`
}

