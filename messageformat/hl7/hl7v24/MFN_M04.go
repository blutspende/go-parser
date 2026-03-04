package hl7v24

// MFN_M04_MF_CDM - Group struct
type MFN_M04_MF_CDM struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	ChargeDescriptionMaster CDM `hl7:"TAG=CDM"`
	Pricing []PRC `hl7:"TAG=PRC;ATR=optional"`
}

// MFN_M04 - Master files notification - Charge description master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFN_M04
type MFN_M04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfCdm []MFN_M04_MF_CDM `hl7:"GROUP"`
}

