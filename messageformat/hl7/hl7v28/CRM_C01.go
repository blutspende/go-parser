package hl7v28

// CRM_C01_PATIENT - Group struct
type CRM_C01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Patient_visit CRM_C01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C01_PATIENT_PATIENT_VISIT - Group struct
type CRM_C01_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// CRM_C01 - Register a patient on a clinical trial
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/CRM_C01
type CRM_C01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Patient []CRM_C01_PATIENT `hl7:"GROUP"`
}

