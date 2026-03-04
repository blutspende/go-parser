package hl7v271

// CRM_C08_PATIENT - Group struct
type CRM_C08_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Patient_visit CRM_C08_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C08_PATIENT_PATIENT_VISIT - Group struct
type CRM_C08_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CRM_C08 - Patient has gone off phase of clinical trial
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/CRM_C08
type CRM_C08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Patient []CRM_C08_PATIENT `hl7:"GROUP"`
}

