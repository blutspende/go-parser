package hl7v231

// CRM_C06_PATIENT - Group struct
type CRM_C06_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
	ClinicalStudyRegistrationSegment CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhaseSegment []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C06 - Cancel patient entering a phase (clerical mistake)
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/CRM_C06
type CRM_C06 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Patient []CRM_C06_PATIENT `hl7:"GROUP"`
}

