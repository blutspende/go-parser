package hl7v26

import "time"

// DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/DG1
type DG1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	DiagnosisCodeDg1 CWE `hl7:"POS=4"`
	DiagnosisDateTime time.Time `hl7:"POS=6"`
	DiagnosisType string `hl7:"POS=7"`
	DiagnosisPriority string `hl7:"POS=16"`
	DiagnosingClinician []XCN `hl7:"POS=17"`
	DiagnosisClassification string `hl7:"POS=18"`
	ConfidentialIndicator string `hl7:"POS=19"`
	AttestationDateTime time.Time `hl7:"POS=20"`
	DiagnosisIdentifier EI `hl7:"POS=21"`
	DiagnosisActionCode string `hl7:"POS=22"`
	ParentDiagnosis EI `hl7:"POS=23"`
	DrgCclValueCode CWE `hl7:"POS=24"`
	DrgGroupingUsage string `hl7:"POS=25"`
	DrgDiagnosisDeterminationStatus string `hl7:"POS=26"`
	PresentOnAdmissionPoaIndicator string `hl7:"POS=27"`
}

