package hl7v24

// ORU_R01_PATIENT_RESULT - Group struct
type ORU_R01_PATIENT_RESULT struct {
	Patient ORU_R01_PATIENT_RESULT_PATIENT `hl7:"GROUP;ATR=optional"`
	OrderObservation []ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION `hl7:"GROUP"`
}

// ORU_R01_PATIENT_RESULT_PATIENT - Group struct
type ORU_R01_PATIENT_RESULT_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit ORU_R01_PATIENT_RESULT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// ORU_R01_PATIENT_RESULT_PATIENT_VISIT - Group struct
type ORU_R01_PATIENT_RESULT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION - Group struct
type ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Observation []ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION `hl7:"GROUP"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION - Group struct
type ORU_R01_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORU_R01 - Unsolicited transmission of an observation message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORU_R01
type ORU_R01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	PatientResult []ORU_R01_PATIENT_RESULT `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

