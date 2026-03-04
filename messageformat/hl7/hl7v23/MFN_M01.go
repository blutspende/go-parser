package hl7v23

// MFN_M01_MF - Group struct
type MFN_M01_MF struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
}

// MFN_M01 - Master file not otherwise specified
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MFN_M01
type MFN_M01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	Mf []MFN_M01_MF `hl7:"GROUP"`
}

