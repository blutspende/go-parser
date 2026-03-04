package hl7v24

// MFR_M09_MF_QUERY - Group struct
type MFR_M09_MF_QUERY struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
}

// MFR_M09 - Master Files Response - Test/observation (categorical) master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFR_M09
type MFR_M09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfQuery []MFR_M09_MF_QUERY `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

