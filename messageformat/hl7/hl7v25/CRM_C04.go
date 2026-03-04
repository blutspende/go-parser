package hl7v25

// CRM_C04_PATIENT - Group struct
type CRM_C04_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C04 - Patient has gone off a clinical trial
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/CRM_C04
type CRM_C04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	Patient []CRM_C04_PATIENT `hl7:"GROUP"`
}

