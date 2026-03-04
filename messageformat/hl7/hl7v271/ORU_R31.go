package hl7v271

// ORU_R31_PATIENT_OBSERVATION - Group struct
type ORU_R31_PATIENT_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// ORU_R31_VISIT - Group struct
type ORU_R31_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// ORU_R31_TIMING_QTY - Group struct
type ORU_R31_TIMING_QTY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORU_R31_OBSERVATION - Group struct
type ORU_R31_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORU_R31 - Unsolicited New Point-Of-Care Observation Message - Search For An Order
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/ORU_R31
type ORU_R31 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation1 []PRT `hl7:"TAG=PRT;ATR=optional"`
	Patient_observation []ORU_R31_PATIENT_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Visit ORU_R31_VISIT `hl7:"GROUP;ATR=optional"`
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation2 []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing_qty []ORU_R31_TIMING_QTY `hl7:"GROUP;ATR=optional"`
	Observation []ORU_R31_OBSERVATION `hl7:"GROUP"`
}

