package hl7v23

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
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RQA_I09_PROCEDURE - Group struct
type RQA_I09_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	Authorization RQA_I09_PROCEDURE_AUTHORIZATION `hl7:"GROUP;ATR=optional"`
}

// RQA_I09_PROCEDURE_AUTHORIZATION - Group struct
type RQA_I09_PROCEDURE_AUTHORIZATION struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQA_I09_OBSERVATION - Group struct
type RQA_I09_OBSERVATION struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Results []RQA_I09_OBSERVATION_RESULTS `hl7:"GROUP;ATR=optional"`
}

// RQA_I09_OBSERVATION_RESULTS - Group struct
type RQA_I09_OBSERVATION_RESULTS struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RQA_I09_VISIT - Group struct
type RQA_I09_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RQA_I09 - Request for modification to an authorization
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RQA_I09
type RQA_I09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ReferralInformationSegment RF1 `hl7:"TAG=RF1;ATR=optional"`
	Authorization RQA_I09_AUTHORIZATION `hl7:"GROUP;ATR=optional"`
	Provider []RQA_I09_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	GuarantorInsurance RQA_I09_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []RQA_I09_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []RQA_I09_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Visit RQA_I09_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

