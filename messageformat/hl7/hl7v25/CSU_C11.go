package hl7v25

// CSU_C11_PATIENT - Group struct
type CSU_C11_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit CSU_C11_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	StudyPhase []CSU_C11_PATIENT_STUDY_PHASE `hl7:"GROUP"`
}

// CSU_C11_PATIENT_VISIT - Group struct
type CSU_C11_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// CSU_C11_PATIENT_STUDY_PHASE - Group struct
type CSU_C11_PATIENT_STUDY_PHASE struct {
	ClinicalStudyPhase CSP `hl7:"TAG=CSP;ATR=optional"`
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
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	TimingQty []CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_OBSERVATION_TIMING_QTY `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX"`
}

// CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_OBSERVATION_TIMING_QTY - Group struct
type CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_OBSERVATION_TIMING_QTY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM - Group struct
type CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	RxAdmin []CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM_RX_ADMIN `hl7:"GROUP"`
}

// CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM_RX_ADMIN - Group struct
type CSU_C11_PATIENT_STUDY_PHASE_STUDY_SCHEDULE_STUDY_PHARM_RX_ADMIN struct {
	PharmacyTreatmentAdministration RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR"`
}

// CSU_C11 - Patient completes a phase of the clinical trial
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/CSU_C11
type CSU_C11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	Patient []CSU_C11_PATIENT `hl7:"GROUP"`
}

