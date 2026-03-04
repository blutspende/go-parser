package hl7v231

// MFN_M02_MF_STAFF - Group struct
type MFN_M02_MF_STAFF struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
	StaffIdentificationSegment STF `hl7:"TAG=STF"`
	PractitionerDetailSegment PRA `hl7:"TAG=PRA;ATR=optional"`
}

// MFN_M02 - Master file - staff practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFN_M02
type MFN_M02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfStaff []MFN_M02_MF_STAFF `hl7:"GROUP"`
}

