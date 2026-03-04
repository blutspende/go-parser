package hl7v26

// MFR_M01_MF_QUERY - Group struct
type MFR_M01_MF_QUERY struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
}

// MFR_M01 - Master files response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/MFR_M01
type MFR_M01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfQuery []MFR_M01_MF_QUERY `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

