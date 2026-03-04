package hl7v251

// MFN_M06_MF_CLIN_STUDY - Group struct
type MFN_M06_MF_CLIN_STUDY struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	ClinicalStudyMaster CM0 `hl7:"TAG=CM0"`
	MfPhaseScheduleDetail []MFN_M06_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFN_M06_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL - Group struct
type MFN_M06_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL struct {
	ClinicalStudyPhaseMaster CM1 `hl7:"TAG=CM1"`
	ClinicalStudyScheduleMaster []CM2 `hl7:"TAG=CM2;ATR=optional"`
}

// MFN_M06 - Master files notification - Clinical study with phases and schedules master file
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/MFN_M06
type MFN_M06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfClinStudy []MFN_M06_MF_CLIN_STUDY `hl7:"GROUP"`
}

