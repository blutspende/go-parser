package hl7v231

// MFN_M04_MF_CDM - Group struct
type MFN_M04_MF_CDM struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
	ChargeDescriptionMasterSegment CDM `hl7:"TAG=CDM"`
	PricingSegment []PRC `hl7:"TAG=PRC;ATR=optional"`
}

// MFN_M04 - Master files charge description
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFN_M04
type MFN_M04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfCdm []MFN_M04_MF_CDM `hl7:"GROUP"`
}

