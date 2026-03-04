package hl7v28

// CQU_I19_PROVIDER_CONTACT - Group struct
type CQU_I19_PROVIDER_CONTACT struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// CQU_I19_PATIENT - Group struct
type CQU_I19_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
}

// CQU_I19_INSURANCE - Group struct
type CQU_I19_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// CQU_I19_APPOINTMENT_HISTORY - Group struct
type CQU_I19_APPOINTMENT_HISTORY struct {
	SchedulingActivityInformation SCH `hl7:"TAG=SCH"`
	Resources []CQU_I19_APPOINTMENT_HISTORY_RESOURCES `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_APPOINTMENT_HISTORY_RESOURCES - Group struct
type CQU_I19_APPOINTMENT_HISTORY_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Resource_detail []CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL - Group struct
type CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL struct {
	Resource_object CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT `hl7:"GROUP"`
	Resource_observation []CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT - Group struct
type CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT struct {
	AppointmentInformation AIS `hl7:"TAG=AIS;ATR=optional"`
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG;ATR=optional"`
	AppointmentInformationLocationResource AIL `hl7:"TAG=AIL;ATR=optional"`
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP;ATR=optional"`
}

// CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBSERVATION - Group struct
type CQU_I19_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CQU_I19_CLINICAL_HISTORY - Group struct
type CQU_I19_CLINICAL_HISTORY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Clinical_history_detail []CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL `hl7:"GROUP;ATR=optional"`
	Role_clinical_history []CQU_I19_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL - Group struct
type CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL struct {
	Clinical_history_object CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT `hl7:"GROUP;ATR=optional"`
	Clinical_history_observation []CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT - Group struct
type CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	Procedures PR1 `hl7:"TAG=PR1;ATR=optional"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	PatientAllergyInformation AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientAdverseReactionInformation IAM `hl7:"TAG=IAM;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	RiskManagementIncident RMI `hl7:"TAG=RMI;ATR=optional"`
	Disability DB1 `hl7:"TAG=DB1;ATR=optional"`
	Diagnosis DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientDeathAndAutopsy PDA `hl7:"TAG=PDA;ATR=optional"`
}

// CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBSERVATION - Group struct
type CQU_I19_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CQU_I19_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY - Group struct
type CQU_I19_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY struct {
	Role_clinical_history_object CQU_I19_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT `hl7:"GROUP;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CQU_I19_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT - Group struct
type CQU_I19_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CQU_I19_PATIENT_VISIT - Group struct
type CQU_I19_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// CQU_I19_MEDICATION_HISTORY - Group struct
type CQU_I19_MEDICATION_HISTORY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Medication_order_detail CQU_I19_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Medication_encoding_detail CQU_I19_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL `hl7:"GROUP;ATR=optional"`
	Medication_administration_detail []CQU_I19_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CQU_I19_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL - Group struct
type CQU_I19_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Medication_order_observation []CQU_I19_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL_MEDICATION_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL_MEDICATION_ORDER_OBSERVATION - Group struct
type CQU_I19_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL_MEDICATION_ORDER_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CQU_I19_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL - Group struct
type CQU_I19_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Medication_encoding_observation []CQU_I19_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL_MEDICATION_ENCODING_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL_MEDICATION_ENCODING_OBSERVATION - Group struct
type CQU_I19_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL_MEDICATION_ENCODING_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CQU_I19_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL - Group struct
type CQU_I19_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL struct {
	PharmacyTreatmentAdministration []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR"`
	Medication_administration_observation []CQU_I19_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL_MEDICATION_ADMINISTRATION_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL_MEDICATION_ADMINISTRATION_OBSERVATION - Group struct
type CQU_I19_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL_MEDICATION_ADMINISTRATION_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CQU_I19_PROBLEM - Group struct
type CQU_I19_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_problem []CQU_I19_PROBLEM_ROLE_PROBLEM `hl7:"GROUP;ATR=optional"`
	Problem_observation []CQU_I19_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_PROBLEM_ROLE_PROBLEM - Group struct
type CQU_I19_PROBLEM_ROLE_PROBLEM struct {
	Role_problem_object CQU_I19_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT `hl7:"GROUP"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CQU_I19_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT - Group struct
type CQU_I19_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CQU_I19_PROBLEM_PROBLEM_OBSERVATION - Group struct
type CQU_I19_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CQU_I19_GOAL - Group struct
type CQU_I19_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_goal []CQU_I19_GOAL_ROLE_GOAL `hl7:"GROUP;ATR=optional"`
	Goal_observation []CQU_I19_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_GOAL_ROLE_GOAL - Group struct
type CQU_I19_GOAL_ROLE_GOAL struct {
	Role_goal_object CQU_I19_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT `hl7:"GROUP"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CQU_I19_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT - Group struct
type CQU_I19_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CQU_I19_GOAL_GOAL_OBSERVATION - Group struct
type CQU_I19_GOAL_GOAL_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CQU_I19_PATHWAY - Group struct
type CQU_I19_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_pathway []CQU_I19_PATHWAY_ROLE_PATHWAY `hl7:"GROUP;ATR=optional"`
	Pathway_observation []CQU_I19_PATHWAY_PATHWAY_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CQU_I19_PATHWAY_ROLE_PATHWAY - Group struct
type CQU_I19_PATHWAY_ROLE_PATHWAY struct {
	Role_pathway_object CQU_I19_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT `hl7:"GROUP"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CQU_I19_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT - Group struct
type CQU_I19_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CQU_I19_PATHWAY_PATHWAY_OBSERVATION - Group struct
type CQU_I19_PATHWAY_PATHWAY_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CQU_I19 - Collaborative Care Referral
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/CQU_I19
type CQU_I19 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	Provider_contact CQU_I19_PROVIDER_CONTACT `hl7:"GROUP;ATR=optional"`
	Patient []CQU_I19_PATIENT `hl7:"GROUP;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Insurance []CQU_I19_INSURANCE `hl7:"GROUP;ATR=optional"`
	Appointment_history []CQU_I19_APPOINTMENT_HISTORY `hl7:"GROUP;ATR=optional"`
	Clinical_history []CQU_I19_CLINICAL_HISTORY `hl7:"GROUP;ATR=optional"`
	Patient_visit []CQU_I19_PATIENT_VISIT `hl7:"GROUP"`
	Medication_history []CQU_I19_MEDICATION_HISTORY `hl7:"GROUP;ATR=optional"`
	Problem []CQU_I19_PROBLEM `hl7:"GROUP;ATR=optional"`
	Goal []CQU_I19_GOAL `hl7:"GROUP;ATR=optional"`
	Pathway []CQU_I19_PATHWAY `hl7:"GROUP;ATR=optional"`
	ClinicalRelationshipSegment []REL `hl7:"TAG=REL;ATR=optional"`
}

