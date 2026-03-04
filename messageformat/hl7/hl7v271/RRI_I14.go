package hl7v271

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
	Procedures PR1 `hl7:"TAG=PR1"`
	AuthorizationContact RRI_I14_PROCEDURE_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
}

// RRI_I14_PROCEDURE_AUTHORIZATION_CONTACT - Group struct
type RRI_I14_PROCEDURE_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RRI_I14_OBSERVATION - Group struct
type RRI_I14_OBSERVATION struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ResultsNotes []RRI_I14_OBSERVATION_RESULTS_NOTES `hl7:"GROUP;ATR=optional"`
}

// RRI_I14_OBSERVATION_RESULTS_NOTES - Group struct
type RRI_I14_OBSERVATION_RESULTS_NOTES struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRI_I14_PATIENT_VISIT - Group struct
type RRI_I14_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RRI_I14 - Cancel Patient Referral Response
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/RRI_I14
type RRI_I14 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA;ATR=optional"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	AuthorizationContact RRI_I14_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
	ProviderContact []RRI_I14_PROVIDER_CONTACT `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []RRI_I14_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []RRI_I14_OBSERVATION `hl7:"GROUP;ATR=optional"`
	PatientVisit RRI_I14_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

