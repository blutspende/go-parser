package hl7v24

// MFK_M07_MF_CLIN_STUDY_SCHEDULE - Group struct
type MFK_M07_MF_CLIN_STUDY_SCHEDULE struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	ClinicalStudyMaster CM0 `hl7:"TAG=CM0"`
	ClinicalStudyScheduleMaster []CM2 `hl7:"TAG=CM2;ATR=optional"`
}

// MFK_M07 - Master file acknowledgment - Clinical study without phases but with schedules master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFK_M07
type MFK_M07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
	MfClinStudySchedule []MFK_M07_MF_CLIN_STUDY_SCHEDULE `hl7:"GROUP"`
}

