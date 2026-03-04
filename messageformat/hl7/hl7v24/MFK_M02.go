package hl7v24

// MFK_M02_MF_STAFF - Group struct
type MFK_M02_MF_STAFF struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	StaffIdentification STF `hl7:"TAG=STF"`
	PractitionerDetail PRA `hl7:"TAG=PRA;ATR=optional"`
	PractitionerOrganizationUnit ORG `hl7:"TAG=ORG;ATR=optional"`
}

// MFK_M02 - Master file acknowledgment - Staff/practitioner master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFK_M02
type MFK_M02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
	MfStaff []MFK_M02_MF_STAFF `hl7:"GROUP"`
}

