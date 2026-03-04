package hl7v24

// MFN_M02_MF_STAFF - Group struct
type MFN_M02_MF_STAFF struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	StaffIdentification STF `hl7:"TAG=STF"`
	PractitionerDetail PRA `hl7:"TAG=PRA;ATR=optional"`
	PractitionerOrganizationUnit ORG `hl7:"TAG=ORG;ATR=optional"`
}

// MFN_M02 - Master files notification - Staff/practitioner master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFN_M02
type MFN_M02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfStaff []MFN_M02_MF_STAFF `hl7:"GROUP"`
}

