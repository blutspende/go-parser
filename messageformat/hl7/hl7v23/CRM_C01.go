package hl7v23

// CRM_C01_PATIENT - Group struct
type CRM_C01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C01 - Register a patient on a clinical trial
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/CRM_C01
type CRM_C01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Patient []CRM_C01_PATIENT `hl7:"GROUP"`
}

