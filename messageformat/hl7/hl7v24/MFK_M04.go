package hl7v24

// MFK_M04_MF_CDM - Group struct
type MFK_M04_MF_CDM struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	ChargeDescriptionMaster CDM `hl7:"TAG=CDM"`
	Pricing []PRC `hl7:"TAG=PRC;ATR=optional"`
}

// MFK_M04 - Master file acknowledgment - Charge description master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFK_M04
type MFK_M04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
	MfCdm []MFK_M04_MF_CDM `hl7:"GROUP"`
}

