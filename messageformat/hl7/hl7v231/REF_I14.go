package hl7v231

// REF_I14_AUTHORIZATION_CONTACT - Group struct
type REF_I14_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I14_PROVIDER - Group struct
type REF_I14_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I14_INSURANCE - Group struct
type REF_I14_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// REF_I14_PROCEDURE - Group struct
type REF_I14_PROCEDURE struct {
	ProceduresSegment PR1 `hl7:"TAG=PR1"`
	AuthorizationContact REF_I14_PROCEDURE_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
}

// REF_I14_PROCEDURE_AUTHORIZATION_CONTACT - Group struct
type REF_I14_PROCEDURE_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I14_OBSERVATION - Group struct
type REF_I14_OBSERVATION struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	ResultsNotes []REF_I14_OBSERVATION_RESULTS_NOTES `hl7:"GROUP;ATR=optional"`
}

// REF_I14_OBSERVATION_RESULTS_NOTES - Group struct
type REF_I14_OBSERVATION_RESULTS_NOTES struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// REF_I14_PATIENT_VISIT - Group struct
type REF_I14_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// REF_I14 - Cancel patient referral
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/REF_I14
type REF_I14 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	AuthorizationContact REF_I14_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
	Provider []REF_I14_PROVIDER `hl7:"GROUP"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	GuarantorSegment []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []REF_I14_INSURANCE `hl7:"GROUP;ATR=optional"`
	AccidentSegment ACC `hl7:"TAG=ACC;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroupSegment []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []REF_I14_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []REF_I14_OBSERVATION `hl7:"GROUP;ATR=optional"`
	PatientVisit REF_I14_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

