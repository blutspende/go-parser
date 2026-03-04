package hl7v24

// OML_O21_PATIENT - Group struct
type OML_O21_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit OML_O21_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OML_O21_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OML_O21_PATIENT_PATIENT_VISIT - Group struct
type OML_O21_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// OML_O21_PATIENT_INSURANCE - Group struct
type OML_O21_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OML_O21_ORDER_GENERAL - Group struct
type OML_O21_ORDER_GENERAL struct {
	Container1 OML_O21_ORDER_GENERAL_CONTAINER_1 `hl7:"GROUP;ATR=optional"`
	Order []OML_O21_ORDER_GENERAL_ORDER `hl7:"GROUP"`
}

// OML_O21_ORDER_GENERAL_CONTAINER_1 - Group struct
type OML_O21_ORDER_GENERAL_CONTAINER_1 struct {
	SpecimenAndContainerDetail SAC `hl7:"TAG=SAC"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// OML_O21_ORDER_GENERAL_ORDER - Group struct
type OML_O21_ORDER_GENERAL_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST - Group struct
type OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	Container2 []OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_CONTAINER_2 `hl7:"GROUP;ATR=optional"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	Observation []OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_OBSERVATION `hl7:"GROUP;ATR=optional"`
	PriorResult []OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT `hl7:"GROUP;ATR=optional"`
}

// OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_CONTAINER_2 - Group struct
type OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_CONTAINER_2 struct {
	SpecimenAndContainerDetail SAC `hl7:"TAG=SAC"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_OBSERVATION - Group struct
type OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT - Group struct
type OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT struct {
	PatientPrior OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_PRIOR `hl7:"GROUP;ATR=optional"`
	PatientVisitPrior OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_VISIT_PRIOR `hl7:"GROUP;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	OrderPrior []OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR `hl7:"GROUP"`
}

// OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_PRIOR - Group struct
type OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_PRIOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
}

// OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_VISIT_PRIOR - Group struct
type OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_PATIENT_VISIT_PRIOR struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR - Group struct
type OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ObservationPrior []OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR `hl7:"GROUP"`
}

// OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR - Group struct
type OML_O21_ORDER_GENERAL_ORDER_OBSERVATION_REQUEST_PRIOR_RESULT_ORDER_PRIOR_OBSERVATION_PRIOR struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OML_O21 - Laboratory order
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/OML_O21
type OML_O21 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OML_O21_PATIENT `hl7:"GROUP;ATR=optional"`
	OrderGeneral []OML_O21_ORDER_GENERAL `hl7:"GROUP"`
}

