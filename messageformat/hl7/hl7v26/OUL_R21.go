package hl7v26

// OUL_R21_PATIENT - Group struct
type OUL_R21_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit OUL_R21_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// OUL_R21_PATIENT_VISIT - Group struct
type OUL_R21_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// OUL_R21_ORDER_OBSERVATION - Group struct
type OUL_R21_ORDER_OBSERVATION struct {
	Container OUL_R21_ORDER_OBSERVATION_CONTAINER `hl7:"GROUP;ATR=optional"`
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	TimingQty []OUL_R21_ORDER_OBSERVATION_TIMING_QTY `hl7:"GROUP;ATR=optional"`
	Observation []OUL_R21_ORDER_OBSERVATION_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// OUL_R21_ORDER_OBSERVATION_CONTAINER - Group struct
type OUL_R21_ORDER_OBSERVATION_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	SubstanceIdentifier SID `hl7:"TAG=SID;ATR=optional"`
}

// OUL_R21_ORDER_OBSERVATION_TIMING_QTY - Group struct
type OUL_R21_ORDER_OBSERVATION_TIMING_QTY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OUL_R21_ORDER_OBSERVATION_OBSERVATION - Group struct
type OUL_R21_ORDER_OBSERVATION_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	SubstanceIdentifier []SID `hl7:"TAG=SID;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OUL_R21 - Unsolicited laboratory observation
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/OUL_R21
type OUL_R21 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OUL_R21_PATIENT `hl7:"GROUP;ATR=optional"`
	OrderObservation []OUL_R21_ORDER_OBSERVATION `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

