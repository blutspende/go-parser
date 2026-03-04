package hl7v231

// CRM_C04_PATIENT - Group struct
type CRM_C04_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
	ClinicalStudyRegistrationSegment CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhaseSegment []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C04 - Patient has gone off a clinical trial
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/CRM_C04
type CRM_C04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Patient []CRM_C04_PATIENT `hl7:"GROUP"`
}

