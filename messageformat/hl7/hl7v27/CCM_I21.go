package hl7v27

// CCM_I21_INSURANCE - Group struct
type CCM_I21_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// CCM_I21_APPOINTMENT_HISTORY - Group struct
type CCM_I21_APPOINTMENT_HISTORY struct {
	SchedulingActivityInformation SCH `hl7:"TAG=SCH"`
	Resources []CCM_I21_APPOINTMENT_HISTORY_RESOURCES `hl7:"GROUP;ATR=optional"`
}

// CCM_I21_APPOINTMENT_HISTORY_RESOURCES - Group struct
type CCM_I21_APPOINTMENT_HISTORY_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Resource_detail []CCM_I21_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL `hl7:"GROUP;ATR=optional"`
}

// CCM_I21_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL - Group struct
type CCM_I21_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL struct {
	AppointmentInformation AIS `hl7:"TAG=AIS"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCM_I21_CLINICAL_HISTORY - Group struct
type CCM_I21_CLINICAL_HISTORY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Clinical_history_detail []CCM_I21_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL `hl7:"GROUP;ATR=optional"`
	Role_clinical_history []CCM_I21_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CCM_I21_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL - Group struct
type CCM_I21_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCM_I21_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY - Group struct
type CCM_I21_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCM_I21_PATIENT_VISITS - Group struct
type CCM_I21_PATIENT_VISITS struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// CCM_I21_MEDICATION_HISTORY - Group struct
type CCM_I21_MEDICATION_HISTORY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Medication_order_detail CCM_I21_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Medication_encoding_detail CCM_I21_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL `hl7:"GROUP;ATR=optional"`
	Medication_administration_detail []CCM_I21_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CCM_I21_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL - Group struct
type CCM_I21_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCM_I21_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL - Group struct
type CCM_I21_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCM_I21_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL - Group struct
type CCM_I21_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL struct {
	PharmacyTreatmentAdministration []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCM_I21_PROBLEM - Group struct
type CCM_I21_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_problem []CCM_I21_PROBLEM_ROLE_PROBLEM `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCM_I21_PROBLEM_ROLE_PROBLEM - Group struct
type CCM_I21_PROBLEM_ROLE_PROBLEM struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCM_I21_GOAL - Group struct
type CCM_I21_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_goal []CCM_I21_GOAL_ROLE_GOAL `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCM_I21_GOAL_ROLE_GOAL - Group struct
type CCM_I21_GOAL_ROLE_GOAL struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCM_I21_PATHWAY - Group struct
type CCM_I21_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_pathway []CCM_I21_PATHWAY_ROLE_PATHWAY `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCM_I21_PATHWAY_ROLE_PATHWAY - Group struct
type CCM_I21_PATHWAY_ROLE_PATHWAY struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCM_I21 - Collaborative Care Message
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/CCM_I21
type CCM_I21 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Insurance []CCM_I21_INSURANCE `hl7:"GROUP;ATR=optional"`
	Appointment_history []CCM_I21_APPOINTMENT_HISTORY `hl7:"GROUP;ATR=optional"`
	Clinical_history []CCM_I21_CLINICAL_HISTORY `hl7:"GROUP;ATR=optional"`
	Patient_visits []CCM_I21_PATIENT_VISITS `hl7:"GROUP"`
	Medication_history []CCM_I21_MEDICATION_HISTORY `hl7:"GROUP;ATR=optional"`
	Problem []CCM_I21_PROBLEM `hl7:"GROUP;ATR=optional"`
	Goal []CCM_I21_GOAL `hl7:"GROUP;ATR=optional"`
	Pathway []CCM_I21_PATHWAY `hl7:"GROUP;ATR=optional"`
	ClinicalRelationshipSegment []REL `hl7:"TAG=REL;ATR=optional"`
}

