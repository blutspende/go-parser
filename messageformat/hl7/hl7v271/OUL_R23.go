package hl7v271

// OUL_R23_PATIENT - Group struct
type OUL_R23_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient_observation []OUL_R23_PATIENT_PATIENT_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Visit OUL_R23_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// OUL_R23_PATIENT_PATIENT_OBSERVATION - Group struct
type OUL_R23_PATIENT_PATIENT_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OUL_R23_PATIENT_VISIT - Group struct
type OUL_R23_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OUL_R23_SPECIMEN - Group struct
type OUL_R23_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	Specimen_observation []OUL_R23_SPECIMEN_SPECIMEN_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Container []OUL_R23_SPECIMEN_CONTAINER `hl7:"GROUP"`
}

// OUL_R23_SPECIMEN_SPECIMEN_OBSERVATION - Group struct
type OUL_R23_SPECIMEN_SPECIMEN_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OUL_R23_SPECIMEN_CONTAINER - Group struct
type OUL_R23_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	InventoryDetail INV `hl7:"TAG=INV;ATR=optional"`
	Order []OUL_R23_SPECIMEN_CONTAINER_ORDER `hl7:"GROUP"`
}

// OUL_R23_SPECIMEN_CONTAINER_ORDER - Group struct
type OUL_R23_SPECIMEN_CONTAINER_ORDER struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing_qty []OUL_R23_SPECIMEN_CONTAINER_ORDER_TIMING_QTY `hl7:"GROUP;ATR=optional"`
	Result []OUL_R23_SPECIMEN_CONTAINER_ORDER_RESULT `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// OUL_R23_SPECIMEN_CONTAINER_ORDER_TIMING_QTY - Group struct
type OUL_R23_SPECIMEN_CONTAINER_ORDER_TIMING_QTY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OUL_R23_SPECIMEN_CONTAINER_ORDER_RESULT - Group struct
type OUL_R23_SPECIMEN_CONTAINER_ORDER_RESULT struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	SubstanceIdentifier []SID `hl7:"TAG=SID;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OUL_R23 - Unsolicited Specimen Container Oriented Observation Message
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/OUL_R23
type OUL_R23 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OUL_R23_PATIENT `hl7:"GROUP;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Specimen []OUL_R23_SPECIMEN `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

