package hl7v28

// OML_O35_PATIENT - Group struct
type OML_O35_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Patient_visit OML_O35_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OML_O35_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OML_O35_PATIENT_PATIENT_VISIT - Group struct
type OML_O35_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OML_O35_PATIENT_INSURANCE - Group struct
type OML_O35_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OML_O35_SPECIMEN - Group struct
type OML_O35_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	Specimen_observation []OML_O35_SPECIMEN_SPECIMEN_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Specimen_container []OML_O35_SPECIMEN_SPECIMEN_CONTAINER `hl7:"GROUP"`
}

// OML_O35_SPECIMEN_SPECIMEN_OBSERVATION - Group struct
type OML_O35_SPECIMEN_SPECIMEN_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	Order []OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER `hl7:"GROUP"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing []OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Observation_request OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_TIMING - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	Observation []OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Prior_result []OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT `hl7:"GROUP;ATR=optional"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_OBSERVATION - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT struct {
	Patient_prior OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_PRIOR `hl7:"GROUP;ATR=optional"`
	Patient_visit_prior OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_VISIT_PRIOR `hl7:"GROUP;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Order_prior []OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR `hl7:"GROUP"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_PRIOR - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_PRIOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_VISIT_PRIOR - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_VISIT_PRIOR struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation1 []PRT `hl7:"TAG=PRT;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation2 []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing_prior []OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_TIMING_PRIOR `hl7:"GROUP;ATR=optional"`
	Observation_prior []OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR `hl7:"GROUP"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_TIMING_PRIOR - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_TIMING_PRIOR struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR - Group struct
type OML_O35_SPECIMEN_SPECIMEN_CONTAINER_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OML_O35 - Laboratory order for multiple orders related to a single container of a sp
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/OML_O35
type OML_O35 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OML_O35_PATIENT `hl7:"GROUP;ATR=optional"`
	Specimen []OML_O35_SPECIMEN `hl7:"GROUP"`
}

