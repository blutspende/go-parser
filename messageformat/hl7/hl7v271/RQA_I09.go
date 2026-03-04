package hl7v271

// RQA_I09_AUTHORIZATION - Group struct
type RQA_I09_AUTHORIZATION struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQA_I09_PROVIDER - Group struct
type RQA_I09_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQA_I09_GUARANTOR_INSURANCE - Group struct
type RQA_I09_GUARANTOR_INSURANCE struct {
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RQA_I09_GUARANTOR_INSURANCE_INSURANCE `hl7:"GROUP"`
}

// RQA_I09_GUARANTOR_INSURANCE_INSURANCE - Group struct
type RQA_I09_GUARANTOR_INSURANCE_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RQA_I09_PROCEDURE - Group struct
type RQA_I09_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	Authorization2 RQA_I09_PROCEDURE_AUTHORIZATION2 `hl7:"GROUP;ATR=optional"`
}

// RQA_I09_PROCEDURE_AUTHORIZATION2 - Group struct
type RQA_I09_PROCEDURE_AUTHORIZATION2 struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQA_I09_OBSERVATION - Group struct
type RQA_I09_OBSERVATION struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Results []RQA_I09_OBSERVATION_RESULTS `hl7:"GROUP;ATR=optional"`
}

// RQA_I09_OBSERVATION_RESULTS - Group struct
type RQA_I09_OBSERVATION_RESULTS struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RQA_I09_VISIT - Group struct
type RQA_I09_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RQA_I09 - Request for modification to an authorization
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/RQA_I09
type RQA_I09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	Authorization RQA_I09_AUTHORIZATION `hl7:"GROUP;ATR=optional"`
	Provider []RQA_I09_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Guarantor_insurance RQA_I09_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []RQA_I09_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []RQA_I09_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Visit RQA_I09_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

