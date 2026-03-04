package hl7v27

// MFN_M07_MF_CLIN_STUDY_SCHEDULE - Group struct
type MFN_M07_MF_CLIN_STUDY_SCHEDULE struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	ClinicalStudyMaster CM0 `hl7:"TAG=CM0"`
	ClinicalStudyScheduleMaster []CM2 `hl7:"TAG=CM2;ATR=optional"`
}

// MFN_M07 - Clinical study without phases but with schedules master file
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/MFN_M07
type MFN_M07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_clin_study_schedule []MFN_M07_MF_CLIN_STUDY_SCHEDULE `hl7:"GROUP"`
}

