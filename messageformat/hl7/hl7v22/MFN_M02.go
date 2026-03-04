package hl7v22

// MFN_M02_MF_STAFF - Group struct
type MFN_M02_MF_STAFF struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	AnyZSegment Zxx `hl7:"TAG=Zxx;ATR=optional"`
}

// MFN_M02 - Staff Practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/MFN_M02
type MFN_M02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_staff []MFN_M02_MF_STAFF `hl7:"GROUP"`
}

