package hl7v231

// ORU_W01_RESPONSE - Group struct
type ORU_W01_RESPONSE struct {
	Patient ORU_W01_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	OrderObservation []ORU_W01_RESPONSE_ORDER_OBSERVATION `hl7:"GROUP"`
}

// ORU_W01_RESPONSE_PATIENT - Group struct
type ORU_W01_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit ORU_W01_RESPONSE_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// ORU_W01_RESPONSE_PATIENT_VISIT - Group struct
type ORU_W01_RESPONSE_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// ORU_W01_RESPONSE_ORDER_OBSERVATION - Group struct
type ORU_W01_RESPONSE_ORDER_OBSERVATION struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []ORU_W01_RESPONSE_ORDER_OBSERVATION_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentificationSegment []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORU_W01_RESPONSE_ORDER_OBSERVATION_OBSERVATION - Group struct
type ORU_W01_RESPONSE_ORDER_OBSERVATION_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORU_W01 - Waveform result, unsolicited transmission of requested information
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ORU_W01
type ORU_W01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Response []ORU_W01_RESPONSE `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

