package hl7v231

// MFN_M07_MF_CLIN_STUDY - Group struct
type MFN_M07_MF_CLIN_STUDY struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
	ClinicalStudyMasterSegment CM0 `hl7:"TAG=CM0"`
	MfPhaseScheduleDetail []MFN_M07_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFN_M07_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL - Group struct
type MFN_M07_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL struct {
	ClinicalStudyPhaseMasterSegment CM1 `hl7:"TAG=CM1"`
	ClinicalStudyScheduleMasterSegment []CM2 `hl7:"TAG=CM2;ATR=optional"`
}

// MFN_M07 - Clinical study without phases but with schedules master file
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFN_M07
type MFN_M07 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfClinStudy []MFN_M07_MF_CLIN_STUDY `hl7:"GROUP"`
}

