package hl7v22

// MFN_M01_MF - Group struct
type MFN_M01_MF struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	AnyZSegment Zxx `hl7:"TAG=Zxx;ATR=optional"`
}

// MFN_M01 - Master files change not otherwise specified (for backwards comp.)
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/MFN_M01
type MFN_M01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf []MFN_M01_MF `hl7:"GROUP"`
}

