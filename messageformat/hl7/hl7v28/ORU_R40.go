package hl7v28

// ORU_R40_PATIENT_RESULT - Group struct
type ORU_R40_PATIENT_RESULT struct {
	Patient ORU_R40_PATIENT_RESULT_PATIENT `hl7:"GROUP;ATR=optional"`
	Order_observation []ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION `hl7:"GROUP"`
}

// ORU_R40_PATIENT_RESULT_PATIENT - Group struct
type ORU_R40_PATIENT_RESULT_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Patient_observation []ORU_R40_PATIENT_RESULT_PATIENT_PATIENT_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Visit ORU_R40_PATIENT_RESULT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// ORU_R40_PATIENT_RESULT_PATIENT_PATIENT_OBSERVATION - Group struct
type ORU_R40_PATIENT_RESULT_PATIENT_PATIENT_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// ORU_R40_PATIENT_RESULT_PATIENT_VISIT - Group struct
type ORU_R40_PATIENT_RESULT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION - Group struct
type ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION struct {
	Common_order ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_COMMON_ORDER `hl7:"GROUP;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing_qty []ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_TIMING_QTY `hl7:"GROUP;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Observation []ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
	Specimen []ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_SPECIMEN `hl7:"GROUP;ATR=optional"`
}

// ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_COMMON_ORDER - Group struct
type ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_COMMON_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Order_document ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_COMMON_ORDER_ORDER_DOCUMENT `hl7:"GROUP;ATR=optional"`
}

// ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_COMMON_ORDER_ORDER_DOCUMENT - Group struct
type ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_COMMON_ORDER_ORDER_DOCUMENT struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	TranscriptionDocumentHeader TXA `hl7:"TAG=TXA"`
}

// ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_TIMING_QTY - Group struct
type ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_TIMING_QTY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION - Group struct
type ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_SPECIMEN - Group struct
type ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	Specimen_observation []ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_SPECIMEN_SPECIMEN_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_SPECIMEN_SPECIMEN_OBSERVATION - Group struct
type ORU_R40_PATIENT_RESULT_ORDER_OBSERVATION_SPECIMEN_SPECIMEN_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// ORU_R40 - Unsolicited Alert Observation Message
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ORU_R40
type ORU_R40 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Patient_result []ORU_R40_PATIENT_RESULT `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

