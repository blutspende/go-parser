package hl7v24

// CRM_C03_PATIENT - Group struct
type CRM_C03_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C03 - Correct/update registration information
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/CRM_C03
type CRM_C03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	Patient []CRM_C03_PATIENT `hl7:"GROUP"`
}

