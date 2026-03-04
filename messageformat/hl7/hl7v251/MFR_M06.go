package hl7v251

// MFR_M06_MF_QUERY - Group struct
type MFR_M06_MF_QUERY struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	ClinicalStudyMaster CM0 `hl7:"TAG=CM0"`
	ClinicalStudyPhaseMaster []CM1 `hl7:"TAG=CM1;ATR=optional"`
	ClinicalStudyScheduleMaster []CM2 `hl7:"TAG=CM2;ATR=optional"`
}

// MFR_M06 - Master Files Response - Clinical study with phases and schedules master file
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/MFR_M06
type MFR_M06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfQuery []MFR_M06_MF_QUERY `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

