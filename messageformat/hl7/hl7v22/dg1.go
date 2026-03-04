package hl7v22

import "time"

// DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/DG1
type DG1 struct {
	SetID                   int       `hl7:"POS=2;ATR=sequence"`
	DiagnosisCodingMethod   string    `hl7:"POS=3"`
	DiagnosisCode           string    `hl7:"POS=4"`
	DiagnosisDescription    string    `hl7:"POS=5"`
	DiagnosisDateTime       time.Time `hl7:"POS=6"`
	DiagnosisDrgType        string    `hl7:"POS=7"`
	MajorDiagnosticCategory CE        `hl7:"POS=8"`
	DiagnosticRelatedGroup  string    `hl7:"POS=9"`
	DrgApprovalIndicator    string    `hl7:"POS=10"`
	DrgGrouperReviewCode    string    `hl7:"POS=11"`
	OutlierType             string    `hl7:"POS=12"`
	OutlierDays             *int      `hl7:"POS=13"`
	OutlierCost             *int      `hl7:"POS=14"`
	GrouperVersionAndType   string    `hl7:"POS=15"`
	DiagnosisDrgPriority    *int      `hl7:"POS=16"`
	DiagnosingClinician     CN_PERSON `hl7:"POS=17"`
}
