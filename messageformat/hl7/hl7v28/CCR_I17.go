package hl7v28

// CCR_I17_PROVIDER_CONTACT - Group struct
type CCR_I17_PROVIDER_CONTACT struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// CCR_I17_CLINICAL_ORDER - Group struct
type CCR_I17_CLINICAL_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Clinical_order_timing []CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Clinical_order_detail []CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL `hl7:"GROUP"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_TIMING - Group struct
type CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL - Group struct
type CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL struct {
	Clinical_order_object CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL_CLINICAL_ORDER_OBJECT `hl7:"GROUP"`
	Clinical_order_observation []CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL_CLINICAL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL_CLINICAL_ORDER_OBJECT - Group struct
type CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL_CLINICAL_ORDER_OBJECT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	Procedures PR1 `hl7:"TAG=PR1;ATR=optional"`
}

// CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL_CLINICAL_ORDER_OBSERVATION - Group struct
type CCR_I17_CLINICAL_ORDER_CLINICAL_ORDER_DETAIL_CLINICAL_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17_PATIENT - Group struct
type CCR_I17_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
}

// CCR_I17_INSURANCE - Group struct
type CCR_I17_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// CCR_I17_APPOINTMENT_HISTORY - Group struct
type CCR_I17_APPOINTMENT_HISTORY struct {
	SchedulingActivityInformation SCH `hl7:"TAG=SCH"`
	Resources []CCR_I17_APPOINTMENT_HISTORY_RESOURCES `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_APPOINTMENT_HISTORY_RESOURCES - Group struct
type CCR_I17_APPOINTMENT_HISTORY_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Resource_detail []CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL - Group struct
type CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL struct {
	Resource_object CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT `hl7:"GROUP"`
	Resource_observation []CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT - Group struct
type CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBJECT struct {
	AppointmentInformation AIS `hl7:"TAG=AIS;ATR=optional"`
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG;ATR=optional"`
	AppointmentInformationLocationResource AIL `hl7:"TAG=AIL;ATR=optional"`
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP;ATR=optional"`
}

// CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBSERVATION - Group struct
type CCR_I17_APPOINTMENT_HISTORY_RESOURCES_RESOURCE_DETAIL_RESOURCE_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17_CLINICAL_HISTORY - Group struct
type CCR_I17_CLINICAL_HISTORY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Clinical_history_detail []CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL `hl7:"GROUP;ATR=optional"`
	Role_clinical_history []CCR_I17_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL - Group struct
type CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL struct {
	Clinical_history_object CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT `hl7:"GROUP"`
	Clinical_history_observation []CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT - Group struct
type CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBJECT struct {
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
}

// CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBSERVATION - Group struct
type CCR_I17_CLINICAL_HISTORY_CLINICAL_HISTORY_DETAIL_CLINICAL_HISTORY_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY - Group struct
type CCR_I17_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY struct {
	Role_clinical_history_object CCR_I17_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT `hl7:"GROUP"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCR_I17_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT - Group struct
type CCR_I17_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_ROLE_CLINICAL_HISTORY_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CCR_I17_PATIENT_VISITS - Group struct
type CCR_I17_PATIENT_VISITS struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// CCR_I17_MEDICATION_HISTORY - Group struct
type CCR_I17_MEDICATION_HISTORY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Medication_order_detail CCR_I17_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Medication_encoding_detail CCR_I17_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL `hl7:"GROUP;ATR=optional"`
	Medication_administration_detail []CCR_I17_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// CCR_I17_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL - Group struct
type CCR_I17_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Medication_order_observation []CCR_I17_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL_MEDICATION_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL_MEDICATION_ORDER_OBSERVATION - Group struct
type CCR_I17_MEDICATION_HISTORY_MEDICATION_ORDER_DETAIL_MEDICATION_ORDER_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL - Group struct
type CCR_I17_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Medication_encoding_observation []CCR_I17_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL_MEDICATION_ENCODING_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL_MEDICATION_ENCODING_OBSERVATION - Group struct
type CCR_I17_MEDICATION_HISTORY_MEDICATION_ENCODING_DETAIL_MEDICATION_ENCODING_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL - Group struct
type CCR_I17_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL struct {
	PharmacyTreatmentAdministration []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR"`
	Medication_administration_observation []CCR_I17_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL_MEDICATION_ADMINISTRATION_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL_MEDICATION_ADMINISTRATION_OBSERVATION - Group struct
type CCR_I17_MEDICATION_HISTORY_MEDICATION_ADMINISTRATION_DETAIL_MEDICATION_ADMINISTRATION_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17_PROBLEM - Group struct
type CCR_I17_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_problem []CCR_I17_PROBLEM_ROLE_PROBLEM `hl7:"GROUP;ATR=optional"`
	Problem_observation []CCR_I17_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_PROBLEM_ROLE_PROBLEM - Group struct
type CCR_I17_PROBLEM_ROLE_PROBLEM struct {
	Role_problem_object CCR_I17_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT `hl7:"GROUP"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCR_I17_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT - Group struct
type CCR_I17_PROBLEM_ROLE_PROBLEM_ROLE_PROBLEM_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CCR_I17_PROBLEM_PROBLEM_OBSERVATION - Group struct
type CCR_I17_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17_GOAL - Group struct
type CCR_I17_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_goal []CCR_I17_GOAL_ROLE_GOAL `hl7:"GROUP;ATR=optional"`
	Goal_observation []CCR_I17_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_GOAL_ROLE_GOAL - Group struct
type CCR_I17_GOAL_ROLE_GOAL struct {
	Role_goal_object CCR_I17_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT `hl7:"GROUP"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCR_I17_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT - Group struct
type CCR_I17_GOAL_ROLE_GOAL_ROLE_GOAL_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CCR_I17_GOAL_GOAL_OBSERVATION - Group struct
type CCR_I17_GOAL_GOAL_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17_PATHWAY - Group struct
type CCR_I17_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Role_pathway []CCR_I17_PATHWAY_ROLE_PATHWAY `hl7:"GROUP;ATR=optional"`
	Pathway_observation []CCR_I17_PATHWAY_PATHWAY_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// CCR_I17_PATHWAY_ROLE_PATHWAY - Group struct
type CCR_I17_PATHWAY_ROLE_PATHWAY struct {
	Role_pathway_object CCR_I17_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT `hl7:"GROUP"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// CCR_I17_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT - Group struct
type CCR_I17_PATHWAY_ROLE_PATHWAY_ROLE_PATHWAY_OBJECT struct {
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
	ProviderData PRD `hl7:"TAG=PRD;ATR=optional"`
}

// CCR_I17_PATHWAY_PATHWAY_OBSERVATION - Group struct
type CCR_I17_PATHWAY_PATHWAY_OBSERVATION struct {
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CCR_I17 - Modify Collaborative Care Referral
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/CCR_I17
type CCR_I17 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	ReferralInformation []RF1 `hl7:"TAG=RF1"`
	Provider_contact []CCR_I17_PROVIDER_CONTACT `hl7:"GROUP"`
	Clinical_order []CCR_I17_CLINICAL_ORDER `hl7:"GROUP;ATR=optional"`
	Patient []CCR_I17_PATIENT `hl7:"GROUP"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Insurance []CCR_I17_INSURANCE `hl7:"GROUP;ATR=optional"`
	Appointment_history []CCR_I17_APPOINTMENT_HISTORY `hl7:"GROUP;ATR=optional"`
	Clinical_history []CCR_I17_CLINICAL_HISTORY `hl7:"GROUP;ATR=optional"`
	Patient_visits []CCR_I17_PATIENT_VISITS `hl7:"GROUP"`
	Medication_history []CCR_I17_MEDICATION_HISTORY `hl7:"GROUP;ATR=optional"`
	Problem []CCR_I17_PROBLEM `hl7:"GROUP;ATR=optional"`
	Goal []CCR_I17_GOAL `hl7:"GROUP;ATR=optional"`
	Pathway []CCR_I17_PATHWAY `hl7:"GROUP;ATR=optional"`
	ClinicalRelationshipSegment []REL `hl7:"TAG=REL;ATR=optional"`
}

