package hl7v231

// RRI_I12_AUTHORIZATION_CONTACT - Group struct
type RRI_I12_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RRI_I12_PROVIDER_CONTACT - Group struct
type RRI_I12_PROVIDER_CONTACT struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RRI_I12_PROCEDURE - Group struct
type RRI_I12_PROCEDURE struct {
	ProceduresSegment PR1 `hl7:"TAG=PR1"`
	AuthorizationContact RRI_I12_PROCEDURE_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
}

// RRI_I12_PROCEDURE_AUTHORIZATION_CONTACT - Group struct
type RRI_I12_PROCEDURE_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RRI_I12_OBSERVATION - Group struct
type RRI_I12_OBSERVATION struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	ResultsNotes []RRI_I12_OBSERVATION_RESULTS_NOTES `hl7:"GROUP;ATR=optional"`
}

// RRI_I12_OBSERVATION_RESULTS_NOTES - Group struct
type RRI_I12_OBSERVATION_RESULTS_NOTES struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRI_I12_PATIENT_VISIT - Group struct
type RRI_I12_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RRI_I12 - Return patient referral
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RRI_I12
type RRI_I12 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA;ATR=optional"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	AuthorizationContact RRI_I12_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
	ProviderContact []RRI_I12_PROVIDER_CONTACT `hl7:"GROUP"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	AccidentSegment ACC `hl7:"TAG=ACC;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroupSegment []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []RRI_I12_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []RRI_I12_OBSERVATION `hl7:"GROUP;ATR=optional"`
	PatientVisit RRI_I12_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

