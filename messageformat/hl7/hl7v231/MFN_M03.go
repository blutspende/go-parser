package hl7v231

// MFN_M03_MF_TEST - Group struct
type MFN_M03_MF_TEST struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE;ATR=optional"`
}

// MFN_M03 - Master file - Test/Observation
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFN_M03
type MFN_M03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	Mf_test []MFN_M03_MF_TEST `hl7:"GROUP"`
}

