package hl7v24

// REF_I14_AUTHORIZATION_CONTACT - Group struct
type REF_I14_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I14_PROVIDER_CONTACT - Group struct
type REF_I14_PROVIDER_CONTACT struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I14_INSURANCE - Group struct
type REF_I14_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// REF_I14_PROCEDURE - Group struct
type REF_I14_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	AuthorizationContact REF_I14_PROCEDURE_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
}

// REF_I14_PROCEDURE_AUTHORIZATION_CONTACT - Group struct
type REF_I14_PROCEDURE_AUTHORIZATION_CONTACT struct {
	AuthorizationInformation AUT `hl7:"TAG=AUT"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
}

// REF_I14_OBSERVATION - Group struct
type REF_I14_OBSERVATION struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ResultsNotes []REF_I14_OBSERVATION_RESULTS_NOTES `hl7:"GROUP;ATR=optional"`
}

// REF_I14_OBSERVATION_RESULTS_NOTES - Group struct
type REF_I14_OBSERVATION_RESULTS_NOTES struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// REF_I14_PATIENT_VISIT - Group struct
type REF_I14_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// REF_I14 - Cancel patient referral
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/REF_I14
type REF_I14 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	ReferralInformation RF1 `hl7:"TAG=RF1;ATR=optional"`
	AuthorizationContact REF_I14_AUTHORIZATION_CONTACT `hl7:"GROUP;ATR=optional"`
	ProviderContact []REF_I14_PROVIDER_CONTACT `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []REF_I14_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Procedure []REF_I14_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Observation []REF_I14_OBSERVATION `hl7:"GROUP;ATR=optional"`
	PatientVisit REF_I14_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

