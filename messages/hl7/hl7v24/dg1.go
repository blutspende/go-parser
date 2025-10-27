package hl7v24

import "time"

// HL7 v2.4 - DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/DG1
type DG1 struct {
	SetID                   int       `hl7:"POS=2;ATR=sequence" json:"SetID ,omitempty"`
	DiagnosisCodingMethod   string    `hl7:"POS=3" json:"DiagnosisCodingMethod,omitempty"`
	DiagnosisCode           CE        `hl7:"POS=4" json:"DiagnosisCode,omitempty"`
	DiagnosisDescription    string    `hl7:"POS=5" json:"DiagnosisDescription,omitempty"`
	DiagnosisDateTime       time.Time `hl7:"POS=6" json:"DiagnosisDateTime,omitempty"`
	DiagnosisType           string    `hl7:"POS=7" json:"DiagnosisType,omitempty"`
	MajorDiagnosticCategory CE        `hl7:"POS=8" json:"MajorDiagnosticCategory,omitempty"`
	DiagnosticRelatedGroup  CE        `hl7:"POS=9" json:"DiagnosticRelatedGroup,omitempty"`
	DRGApprovalIndicator    string    `hl7:"POS=10" json:"DRGApprovalIndicator,omitempty"`
	DRGGrouperReviewCode    string    `hl7:"POS=11" json:"DRGGrouperReviewCode,omitempty"`
	OutlierType             CE        `hl7:"POS=12" json:"OutlierType,omitempty"`
	OutlierDays             *int      `hl7:"POS=13" json:"OutlierDays,omitempty"`
	OutlierCost             CP        `hl7:"POS=14" json:"OutlierCost,omitempty"`
	GrouperVersionAndType   string    `hl7:"POS=15" json:"GrouperVersionAndType,omitempty"`
	DiagnosisPriority       *int      `hl7:"POS=16" json:"DiagnosisPriority,omitempty"`
	DiagnosingClinician     []XCN     `hl7:"POS=17" json:"DiagnosingClinician,omitempty"`
	DiagnosisClassification string    `hl7:"POS=18" json:"DiagnosisClassification,omitempty"`
	ConfidentialIndicator   string    `hl7:"POS=19" json:"ConfidentialIndicator,omitempty"`
	AttestationDateTime     time.Time `hl7:"POS=20" json:"AttestationDateTime,omitempty"`
}
