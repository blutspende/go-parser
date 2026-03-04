package hl7v271

// OMG_O19_PATIENT - Group struct
type OMG_O19_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Patient_visit OMG_O19_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OMG_O19_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OMG_O19_PATIENT_PATIENT_VISIT - Group struct
type OMG_O19_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OMG_O19_PATIENT_INSURANCE - Group struct
type OMG_O19_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OMG_O19_ORDER - Group struct
type OMG_O19_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation1 []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing []OMG_O19_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation2 []PRT `hl7:"TAG=PRT;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	Observation []OMG_O19_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Specimen []OMG_O19_ORDER_SPECIMEN `hl7:"GROUP;ATR=optional"`
	Prior_result []OMG_O19_ORDER_PRIOR_RESULT `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OMG_O19_ORDER_TIMING - Group struct
type OMG_O19_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OMG_O19_ORDER_OBSERVATION - Group struct
type OMG_O19_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMG_O19_ORDER_SPECIMEN - Group struct
type OMG_O19_ORDER_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	Specimen_observation []OMG_O19_ORDER_SPECIMEN_SPECIMEN_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Container []OMG_O19_ORDER_SPECIMEN_CONTAINER `hl7:"GROUP;ATR=optional"`
}

// OMG_O19_ORDER_SPECIMEN_SPECIMEN_OBSERVATION - Group struct
type OMG_O19_ORDER_SPECIMEN_SPECIMEN_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OMG_O19_ORDER_SPECIMEN_CONTAINER - Group struct
type OMG_O19_ORDER_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	Container_observation []OMG_O19_ORDER_SPECIMEN_CONTAINER_CONTAINER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// OMG_O19_ORDER_SPECIMEN_CONTAINER_CONTAINER_OBSERVATION - Group struct
type OMG_O19_ORDER_SPECIMEN_CONTAINER_CONTAINER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OMG_O19_ORDER_PRIOR_RESULT - Group struct
type OMG_O19_ORDER_PRIOR_RESULT struct {
	Patient_prior OMG_O19_ORDER_PRIOR_RESULT_PATIENT_PRIOR `hl7:"GROUP;ATR=optional"`
	Patient_visit_prior OMG_O19_ORDER_PRIOR_RESULT_PATIENT_VISIT_PRIOR `hl7:"GROUP;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Order_prior []OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR `hl7:"GROUP"`
}

// OMG_O19_ORDER_PRIOR_RESULT_PATIENT_PRIOR - Group struct
type OMG_O19_ORDER_PRIOR_RESULT_PATIENT_PRIOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OMG_O19_ORDER_PRIOR_RESULT_PATIENT_VISIT_PRIOR - Group struct
type OMG_O19_ORDER_PRIOR_RESULT_PATIENT_VISIT_PRIOR struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR - Group struct
type OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	Timing_prior []OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR_TIMING_PRIOR `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Observation_prior []OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR `hl7:"GROUP"`
}

// OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR_TIMING_PRIOR - Group struct
type OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR_TIMING_PRIOR struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR - Group struct
type OMG_O19_ORDER_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMG_O19 - General clinical order
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/OMG_O19
type OMG_O19 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OMG_O19_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []OMG_O19_ORDER `hl7:"GROUP"`
}

