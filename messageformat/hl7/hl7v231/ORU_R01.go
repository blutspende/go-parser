package hl7v231

// ORU_R01_PATIENT_RESULT - Group struct
type ORU_R01_PATIENT_RESULT struct {
	Patient ORU_R01_PATIENT_RESULT_PATIENT `hl7:"GROUP;ATR=optional"`
	OrderObservation []ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION `hl7:"GROUP"`
}

// ORU_R01_PATIENT_RESULT_PATIENT - Group struct
type ORU_R01_PATIENT_RESULT_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit ORU_R01_PATIENT_RESULT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// ORU_R01_PATIENT_RESULT_PATIENT_VISIT - Group struct
type ORU_R01_PATIENT_RESULT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION - Group struct
type ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentificationSegment []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION - Group struct
type ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORU_R01 - Unsolicited transmission of an observation message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ORU_R01
type ORU_R01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	PatientResult []ORU_R01_PATIENT_RESULT `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

