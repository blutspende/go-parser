package hl7v23

// REF_I12_AUTHORIZATION_CONTACT - Group struct
type REF_I12_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I12_PROVIDER_CONTACT - Group struct
type REF_I12_PROVIDER_CONTACT struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I12_INSURANCE - Group struct
type REF_I12_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// REF_I12_PROCEDURE - Group struct
type REF_I12_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	AuthorizationContact REF_I12_PROCEDURE_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
}

// REF_I12_PROCEDURE_AUTHORIZATION_CONTACT - Group struct
type REF_I12_PROCEDURE_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I12_OBSERVATION - Group struct
type REF_I12_OBSERVATION struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	ResultsNotes []REF_I12_OBSERVATION_RESULTS_NOTES `hl7:"GROUP;ATR=optional"`
}

// REF_I12_OBSERVATION_RESULTS_NOTES - Group struct
type REF_I12_OBSERVATION_RESULTS_NOTES struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// REF_I12_PATIENT_VISIT - Group struct
type REF_I12_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// REF_I12 - Patient referral
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/REF_I12
type REF_I12 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ReferralInformationSegment RF1 `hl7:"TAG=RF1;ATR=optional"`
	AuthorizationContact REF_I12_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
	ProviderContact []REF_I12_PROVIDER_CONTACT `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []REF_I12_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []REF_I12_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []REF_I12_OBSERVATION `hl7:"GROUP;ATR=optional"`
	PatientVisit REF_I12_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

