package hl7v24

// MFK_M06_MF_CLIN_STUDY - Group struct
type MFK_M06_MF_CLIN_STUDY struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	ClinicalStudyMaster CM0 `hl7:"TAG=CM0"`
	MfPhaseScheduleDetail []MFK_M06_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFK_M06_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL - Group struct
type MFK_M06_MF_CLIN_STUDY_MF_PHASE_SCHEDULE_DETAIL struct {
	ClinicalStudyPhaseMaster CM1 `hl7:"TAG=CM1"`
	ClinicalStudyScheduleMaster []CM2 `hl7:"TAG=CM2;ATR=optional"`
}

// MFK_M06 - Master file acknowledgment - Clinical study with phases and schedules master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFK_M06
type MFK_M06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
	MfClinStudy []MFK_M06_MF_CLIN_STUDY `hl7:"GROUP"`
}

