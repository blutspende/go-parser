package hl7v25

// CRM_C06_PATIENT - Group struct
type CRM_C06_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// CRM_C06 - Cancel patient entering a phase (clerical mistake)
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/CRM_C06
type CRM_C06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	Patient []CRM_C06_PATIENT `hl7:"GROUP"`
}

