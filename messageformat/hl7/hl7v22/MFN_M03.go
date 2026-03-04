package hl7v22

// MFN_M03_MF_TEST - Group struct
type MFN_M03_MF_TEST struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	AnyZSegment Zxx `hl7:"TAG=Zxx;ATR=optional"`
}

// MFN_M03 - Test/Observation
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/MFN_M03
type MFN_M03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_test []MFN_M03_MF_TEST `hl7:"GROUP"`
}

