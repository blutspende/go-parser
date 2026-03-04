package hl7v271

// CCI_I22_INSURANCE - Group struct
type CCI_I22_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// CCI_I22_APPOINTMENT_HISTORY - Group struct
type CCI_I22_APPOINTMENT_HISTORY struct {
	SchedulingActivityInformation SCH `hl7:"TAG=SCH"`
	Resources []CCI_I22_APPOINTMENT_HISTORY_RESOURCES `hl7:"GROUP;ATR=optional"`
}

// CCI_I22_APPOINTMENT_HISTORY_RESOURCES - Group struct
type CCI_I22_APPOINTMENT_HISTORY_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Resource_detail []CCI_I22_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL `hl7:"GROUP;ATR=optional"`
}

// CCI_I22_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL - Group struct
type CCI_I22_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL struct {
	Resource_object CCI_I22_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCI_I22_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT - Group struct
type CCI_I22_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT struct {
	AppointmentInformation AIS `hl7:"TAG=AIS;ATR=optional"`
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG;ATR=optional"`
	AppointmentInformationLocationResource AIL `hl7:"TAG=AIL;ATR=optional"`
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP;ATR=optional"`
}

// CCI_I22_CLINICAL_HISTORY - Group struct
type CCI_I22_CLINICAL_HISTORY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Clinical_history_detail []CCI_I22_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL `hl7:"GROUP;ATR=optional"`
	Role_clinical_history []CCI_I22_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CCI_I22_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL - Group struct
type CCI_I22_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL struct {
	Clinical_history_object CCI_I22_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCI_I22_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT - Group struct
type CCI_I22_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT struct {
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

// CCI_I22_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY - Group struct
type CCI_I22_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY struct {
	Role_clinical_history_object CCI_I22_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT `hl7:"GROUP;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCI_I22_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT - Group struct
type CCI_I22_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CCI_I22_PATIENT_VISIT - Group struct
type CCI_I22_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// CCI_I22_MEDICATION_HISTORY - Group struct
type CCI_I22_MEDICATION_HISTORY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Medication_order_detail CCI_I22_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Medication_encoding_detail CCI_I22_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL `hl7:"GROUP;ATR=optional"`
	Medication_administration_detail []CCI_I22_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CCI_I22_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL - Group struct
type CCI_I22_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCI_I22_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL - Group struct
type CCI_I22_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCI_I22_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL - Group struct
type CCI_I22_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL struct {
	PharmacyTreatmentAdministration []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCI_I22_PROBLEM - Group struct
type CCI_I22_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_problem []CCI_I22_PROBLEM_ROLE_PROBLEM `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCI_I22_PROBLEM_ROLE_PROBLEM - Group struct
type CCI_I22_PROBLEM_ROLE_PROBLEM struct {
	Role_problem_object CCI_I22_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT `hl7:"GROUP;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCI_I22_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT - Group struct
type CCI_I22_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CCI_I22_GOAL - Group struct
type CCI_I22_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_goal []CCI_I22_GOAL_ROLE_GOAL `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCI_I22_GOAL_ROLE_GOAL - Group struct
type CCI_I22_GOAL_ROLE_GOAL struct {
	Role_goal_object CCI_I22_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT `hl7:"GROUP;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCI_I22_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT - Group struct
type CCI_I22_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CCI_I22_PATHWAY - Group struct
type CCI_I22_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_pathway []CCI_I22_PATHWAY_ROLE_PATHWAY `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// CCI_I22_PATHWAY_ROLE_PATHWAY - Group struct
type CCI_I22_PATHWAY_ROLE_PATHWAY struct {
	Role_pathway_object CCI_I22_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT `hl7:"GROUP;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCI_I22_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT - Group struct
type CCI_I22_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CCI_I22 - Collaborative Care Information
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/CCI_I22
type CCI_I22 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Insurance []CCI_I22_INSURANCE `hl7:"GROUP;ATR=optional"`
	Appointment_history []CCI_I22_APPOINTMENT_HISTORY `hl7:"GROUP;ATR=optional"`
	Clinical_history []CCI_I22_CLINICAL_HISTORY `hl7:"GROUP;ATR=optional"`
	Patient_visit []CCI_I22_PATIENT_VISIT `hl7:"GROUP"`
	Medication_history []CCI_I22_MEDICATION_HISTORY `hl7:"GROUP;ATR=optional"`
	Problem []CCI_I22_PROBLEM `hl7:"GROUP;ATR=optional"`
	Goal []CCI_I22_GOAL `hl7:"GROUP;ATR=optional"`
	Pathway []CCI_I22_PATHWAY `hl7:"GROUP;ATR=optional"`
	ClinicalRelationshipSegment []REL `hl7:"TAG=REL;ATR=optional"`
}

