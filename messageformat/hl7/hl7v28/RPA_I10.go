package hl7v28

// RPA_I10_AUTHORIZATION_1 - Group struct
type RPA_I10_AUTHORIZATION_1 struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPA_I10_PROVIDER - Group struct
type RPA_I10_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPA_I10_INSURANCE - Group struct
type RPA_I10_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RPA_I10_PROCEDURE - Group struct
type RPA_I10_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	Authorization2 RPA_I10_PROCEDURE_AUTHORIZATION_2 `hl7:"GROUP;ATR=optional"`
}

// RPA_I10_PROCEDURE_AUTHORIZATION_2 - Group struct
type RPA_I10_PROCEDURE_AUTHORIZATION_2 struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPA_I10_OBSERVATION - Group struct
type RPA_I10_OBSERVATION struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Results []RPA_I10_OBSERVATION_RESULTS `hl7:"GROUP;ATR=optional"`
}

// RPA_I10_OBSERVATION_RESULTS - Group struct
type RPA_I10_OBSERVATION_RESULTS struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RPA_I10_VISIT - Group struct
type RPA_I10_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RPA_I10 - Return Resubmission of an Authorization
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RPA_I10
type RPA_I10 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	Authorization1 RPA_I10_AUTHORIZATION_1 `hl7:"GROUP;ATR=optional"`
	Provider []RPA_I10_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RPA_I10_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []RPA_I10_PROCEDURE `hl7:"GROUP"`
	Observation []RPA_I10_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Visit RPA_I10_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

