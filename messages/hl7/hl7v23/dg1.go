package hl7v23

import "time"

// DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/DG1
type DG1 struct {
	SetID                   int       `hl7:"1, sequence" json:"setID,omitempty"`
	DiagnosisCodingMethod   string    `hl7:"2" json:"diagnosisCodingMethod,omitempty"`
	DiagnosisCode           CE        `hl7:"3" json:"diagnosisCode,omitempty"`
	DiagnosisDescription    string    `hl7:"4" json:"diagnosisDescription,omitempty"`
	DiagnosisDateTime       time.Time `hl7:"5,longdate" json:"diagnosisDateTime,omitempty"`
	DiagnosisType           string    `hl7:"6" json:"diagnosisType,omitempty"`
	MajorDiagnosticCategory CE        `hl7:"7" json:"majorDiagnosticCategory,omitempty"`
	DiagnosticRelatedGroup  CE        `hl7:"8" json:"diagnosticRelatedGroup,omitempty"`
	DRGApprovalIndicator    string    `hl7:"9" json:"drgApprovalIndicator,omitempty"`
	DRGGrouperReviewCode    string    `hl7:"10" json:"drgGrouperReviewCode,omitempty"`
	OutlierType             CE        `hl7:"11" json:"outlierType,omitempty"`
	OutlierDays             int       `hl7:"12" json:"outlierDays,omitempty"`
	OutlierCost             CP        `hl7:"13" json:"outlierCost,omitempty"`
	GrouperVersionAndType   string    `hl7:"14" json:"grouperVersionAndType,omitempty"`
	DiagnosisPriority       int       `hl7:"15" json:"diagnosisPriority,omitempty"`
	DiagnosingClinician     []XCN     `hl7:"16" json:"diagnosingClinician,omitempty"`
	DiagnosisClassification string    `hl7:"17" json:"diagnosisClassification,omitempty"`
	ConfidentialIndicator   string    `hl7:"18" json:"confidentialIndicator,omitempty"`
	AttestationDateTime     time.Time `hl7:"19,longdate" json:"attestationDateTime,omitempty"`
}
