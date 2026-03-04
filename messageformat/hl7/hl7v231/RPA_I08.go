package hl7v231

// RPA_I08_AUTHORIZATION - Group struct
type RPA_I08_AUTHORIZATION struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPA_I08_PROVIDER - Group struct
type RPA_I08_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPA_I08_INSURANCE - Group struct
type RPA_I08_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RPA_I08_PROCEDURE - Group struct
type RPA_I08_PROCEDURE struct {
	ProceduresSegment PR1 `hl7:"TAG=PR1"`
	Authorization RPA_I08_PROCEDURE_AUTHORIZATION `hl7:"GROUP;ATR=optional"`
}

// RPA_I08_PROCEDURE_AUTHORIZATION - Group struct
type RPA_I08_PROCEDURE_AUTHORIZATION struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPA_I08_OBSERVATION - Group struct
type RPA_I08_OBSERVATION struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Results []RPA_I08_OBSERVATION_RESULTS `hl7:"GROUP;ATR=optional"`
}

// RPA_I08_OBSERVATION_RESULTS - Group struct
type RPA_I08_OBSERVATION_RESULTS struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RPA_I08_VISIT - Group struct
type RPA_I08_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RPA_I08 - Return for treatment authorization information
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RPA_I08
type RPA_I08 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	Authorization RPA_I08_AUTHORIZATION `hl7:"GROUP;ATR=optional"`
	Provider []RPA_I08_PROVIDER `hl7:"GROUP"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	GuarantorSegment []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RPA_I08_INSURANCE `hl7:"GROUP;ATR=optional"`
	AccidentSegment ACC `hl7:"TAG=ACC;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroupSegment []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []RPA_I08_PROCEDURE `hl7:"GROUP"`
	Observation []RPA_I08_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Visit RPA_I08_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

