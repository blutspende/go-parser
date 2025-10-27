package hl7v23

import "time"

// DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/DG1
type DG1 struct {
	SetID                   int       `hl7:"POS=2;ATR= sequence" json:"setID,omitempty"`
	DiagnosisCodingMethod   string    `hl7:"POS=3" json:"diagnosisCodingMethod,omitempty"`
	DiagnosisCode           CE        `hl7:"POS=4" json:"diagnosisCode,omitempty"`
	DiagnosisDescription    string    `hl7:"POS=5" json:"diagnosisDescription,omitempty"`
	DiagnosisDateTime       time.Time `hl7:"POS=6" json:"diagnosisDateTime,omitempty"`
	DiagnosisType           string    `hl7:"POS=7" json:"diagnosisType,omitempty"`
	MajorDiagnosticCategory CE        `hl7:"POS=8" json:"majorDiagnosticCategory,omitempty"`
	DiagnosticRelatedGroup  CE        `hl7:"POS=9" json:"diagnosticRelatedGroup,omitempty"`
	DRGApprovalIndicator    string    `hl7:"POS=10" json:"drgApprovalIndicator,omitempty"`
	DRGGrouperReviewCode    string    `hl7:"POS=11" json:"drgGrouperReviewCode,omitempty"`
	OutlierType             CE        `hl7:"POS=12" json:"outlierType,omitempty"`
	OutlierDays             *int      `hl7:"POS=13" json:"outlierDays,omitempty"`
	OutlierCost             CP        `hl7:"POS=14" json:"outlierCost,omitempty"`
	GrouperVersionAndType   string    `hl7:"POS=15" json:"grouperVersionAndType,omitempty"`
	DiagnosisPriority       *int      `hl7:"POS=16" json:"diagnosisPriority,omitempty"`
	DiagnosingClinician     []XCN     `hl7:"POS=17" json:"diagnosingClinician,omitempty"`
	DiagnosisClassification string    `hl7:"POS=18" json:"diagnosisClassification,omitempty"`
	ConfidentialIndicator   string    `hl7:"POS=19" json:"confidentialIndicator,omitempty"`
	AttestationDateTime     time.Time `hl7:"POS=20" json:"attestationDateTime,omitempty"`
}
