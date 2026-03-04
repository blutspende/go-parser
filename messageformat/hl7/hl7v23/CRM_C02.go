package hl7v23

// CRM_C02_PATIENT - Group struct
type CRM_C02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C02 - Cancel a patient registration on clinical trial
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/CRM_C02
type CRM_C02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Patient []CRM_C02_PATIENT `hl7:"GROUP"`
}

