package hl7v231

// RRI_I14_AUTHORIZATION_CONTACT - Group struct
type RRI_I14_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RRI_I14_PROVIDER_CONTACT - Group struct
type RRI_I14_PROVIDER_CONTACT struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RRI_I14_PROCEDURE - Group struct
type RRI_I14_PROCEDURE struct {
	ProceduresSegment PR1 `hl7:"TAG=PR1"`
	AuthorizationContact RRI_I14_PROCEDURE_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
}

// RRI_I14_PROCEDURE_AUTHORIZATION_CONTACT - Group struct
type RRI_I14_PROCEDURE_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RRI_I14_OBSERVATION - Group struct
type RRI_I14_OBSERVATION struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	ResultsNotes []RRI_I14_OBSERVATION_RESULTS_NOTES `hl7:"GROUP;ATR=optional"`
}

// RRI_I14_OBSERVATION_RESULTS_NOTES - Group struct
type RRI_I14_OBSERVATION_RESULTS_NOTES struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRI_I14_PATIENT_VISIT - Group struct
type RRI_I14_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RRI_I14 - Return cancel patient referral
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RRI_I14
type RRI_I14 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA;ATR=optional"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	AuthorizationContact RRI_I14_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
	ProviderContact []RRI_I14_PROVIDER_CONTACT `hl7:"GROUP"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	AccidentSegment ACC `hl7:"TAG=ACC;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroupSegment []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []RRI_I14_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []RRI_I14_OBSERVATION `hl7:"GROUP;ATR=optional"`
	PatientVisit RRI_I14_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

