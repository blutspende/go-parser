package hl7v23

// MFR_M09_MF_QUERY - Group struct
type MFR_M09_MF_QUERY struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
}

// MFR_M09 - Master Files Response - Test/Observation
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MFR_M09
type MFR_M09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfQuery []MFR_M09_MF_QUERY `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

