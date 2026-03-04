package hl7v231

// RQA_I11_AUTHORIZATION - Group struct
type RQA_I11_AUTHORIZATION struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQA_I11_PROVIDER - Group struct
type RQA_I11_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQA_I11_GUARANTOR_INSURANCE - Group struct
type RQA_I11_GUARANTOR_INSURANCE struct {
	GuarantorSegment []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RQA_I11_GUARANTOR_INSURANCE_INSURANCE `hl7:"GROUP"`
}

// RQA_I11_GUARANTOR_INSURANCE_INSURANCE - Group struct
type RQA_I11_GUARANTOR_INSURANCE_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RQA_I11_PROCEDURE - Group struct
type RQA_I11_PROCEDURE struct {
	ProceduresSegment PR1 `hl7:"TAG=PR1"`
	Authorization RQA_I11_PROCEDURE_AUTHORIZATION `hl7:"GROUP;ATR=optional"`
}

// RQA_I11_PROCEDURE_AUTHORIZATION - Group struct
type RQA_I11_PROCEDURE_AUTHORIZATION struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQA_I11_OBSERVATION - Group struct
type RQA_I11_OBSERVATION struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Results []RQA_I11_OBSERVATION_RESULTS `hl7:"GROUP;ATR=optional"`
}

// RQA_I11_OBSERVATION_RESULTS - Group struct
type RQA_I11_OBSERVATION_RESULTS struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RQA_I11_VISIT - Group struct
type RQA_I11_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RQA_I11 - Request for cancellation of an authorization
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RQA_I11
type RQA_I11 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	Authorization RQA_I11_AUTHORIZATION `hl7:"GROUP;ATR=optional"`
	Provider []RQA_I11_PROVIDER `hl7:"GROUP"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	GuarantorInsurance RQA_I11_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	AccidentSegment ACC `hl7:"TAG=ACC;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroupSegment []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []RQA_I11_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []RQA_I11_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Visit RQA_I11_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

