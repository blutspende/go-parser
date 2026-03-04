package hl7v26

// MFR_M05_MF_QUERY - Group struct
type MFR_M05_MF_QUERY struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	LocationIdentification LOC `hl7:"TAG=LOC"`
	LocationCharacteristic []LCH `hl7:"TAG=LCH;ATR=optional"`
	LocationRelationship []LRL `hl7:"TAG=LRL;ATR=optional"`
	MfLocation []MFR_M05_MF_QUERY_MF_LOCATION `hl7:"GROUP"`
}

// MFR_M05_MF_QUERY_MF_LOCATION - Group struct
type MFR_M05_MF_QUERY_MF_LOCATION struct {
	LocationDepartment LDP `hl7:"TAG=LDP"`
	LocationCharacteristic []LCH `hl7:"TAG=LCH;ATR=optional"`
	LocationChargeCode []LCC `hl7:"TAG=LCC;ATR=optional"`
}

// MFR_M05 - Master files response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/MFR_M05
type MFR_M05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfQuery []MFR_M05_MF_QUERY `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

