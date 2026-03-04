package hl7v231

// CSU_C11_PATIENT - Group struct
type CSU_C11_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit CSU_C11_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	ClinicalStudyRegistrationSegment CSR `hl7:"TAG=CSR"`
	StudyPhase []CSU_C11_PATIENT_STUDY_PHASE `hl7:"GROUP"`
}

// CSU_C11_PATIENT_VISIT - Group struct
type CSU_C11_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// CSU_C11_PATIENT_STUDY_PHASE - Group struct
type CSU_C11_PATIENT_STUDY_PHASE struct {
	ClinicalStudyPhaseSegment CSP `hl7:"TAG=CSP;ATR=optional"`
	StudySchedule []CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE `hl7:"GROUP"`
}

// CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE - Group struct
type CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE struct {
	ClinicalStudyDataScheduleSegment CSS `hl7:"TAG=CSS;ATR=optional"`
	StudyObservation []CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_OBSERVATION `hl7:"GROUP"`
	StudyPharm []CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM `hl7:"GROUP"`
}

// CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_OBSERVATION - Group struct
type CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_OBSERVATION struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX"`
}

// CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM - Group struct
type CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	RxAdmin []CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM_RX_ADMIN `hl7:"GROUP"`
}

// CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM_RX_ADMIN - Group struct
type CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM_RX_ADMIN struct {
	PharmacyTreatmentAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRouteSegment RXR `hl7:"TAG=RXR"`
}

// CSU_C11 - Patient completes a phase of the clinical trial
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/CSU_C11
type CSU_C11 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Patient []CSU_C11_PATIENT `hl7:"GROUP"`
}

