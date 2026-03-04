package hl7v231

// MFR_M08_MF_QUERY - Group struct
type MFR_M08_MF_QUERY struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
}

// MFR_M08 - Master Files Response - Test/observation (Numeric)
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFR_M08
type MFR_M08 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfQuery []MFR_M08_MF_QUERY `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

