package hl7v231

// CRM_C07_PATIENT - Group struct
type CRM_C07_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
	ClinicalStudyRegistrationSegment CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhaseSegment []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C07 - Correct/update phase information
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/CRM_C07
type CRM_C07 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Patient []CRM_C07_PATIENT `hl7:"GROUP"`
}

