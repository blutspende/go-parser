package hl7v251

import "time"

// DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/DG1
type DG1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	DiagnosisCodingMethod string `hl7:"POS=3"`
	DiagnosisCodeDg1 CE `hl7:"POS=4"`
	DiagnosisDescription string `hl7:"POS=5"`
	DiagnosisDateTime time.Time `hl7:"POS=6"`
	DiagnosisType string `hl7:"POS=7"`
	MajorDiagnosticCategory CE `hl7:"POS=8"`
	DiagnosticRelatedGroup CE `hl7:"POS=9"`
	DrgApprovalIndicator string `hl7:"POS=10"`
	DrgGrouperReviewCode string `hl7:"POS=11"`
	OutlierType CE `hl7:"POS=12"`
	OutlierDays *int `hl7:"POS=13"`
	OutlierCost CP `hl7:"POS=14"`
	GrouperVersionAndType string `hl7:"POS=15"`
	DiagnosisPriority string `hl7:"POS=16"`
	DiagnosingClinician []XCN `hl7:"POS=17"`
	DiagnosisClassification string `hl7:"POS=18"`
	ConfidentialIndicator string `hl7:"POS=19"`
	AttestationDateTime time.Time `hl7:"POS=20"`
	DiagnosisIdentifier EI `hl7:"POS=21"`
	DiagnosisActionCode string `hl7:"POS=22"`
}

