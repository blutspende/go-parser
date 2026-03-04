package hl7v27

// CRM_C02_PATIENT - Group struct
type CRM_C02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Patient_visit CRM_C02_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C02_PATIENT_PATIENT_VISIT - Group struct
type CRM_C02_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CRM_C02 - Cancel a patient registration on clinical trial (for clerical mistakes onl
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/CRM_C02
type CRM_C02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Patient []CRM_C02_PATIENT `hl7:"GROUP"`
}

