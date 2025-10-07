package hl7v24

import "time"

// HL7 v2.4 - DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/DG1
type DG1 struct {
	SetID                   int       `hl7:"1, sequence" json:"SetID ,omitempty"`
	DiagnosisCodingMethod   string    `hl7:"2" json:"DiagnosisCodingMethod,omitempty"`
	DiagnosisCode           CE        `hl7:"3" json:"DiagnosisCode,omitempty"`
	DiagnosisDescription    string    `hl7:"4" json:"DiagnosisDescription,omitempty"`
	DiagnosisDateTime       time.Time `hl7:"5,longdate" json:"DiagnosisDateTime,omitempty"`
	DiagnosisType           string    `hl7:"6" json:"DiagnosisType,omitempty"`
	MajorDiagnosticCategory CE        `hl7:"7" json:"MajorDiagnosticCategory,omitempty"`
	DiagnosticRelatedGroup  CE        `hl7:"8" json:"DiagnosticRelatedGroup,omitempty"`
	DRGApprovalIndicator    string    `hl7:"9" json:"DRGApprovalIndicator,omitempty"`
	DRGGrouperReviewCode    string    `hl7:"10" json:"DRGGrouperReviewCode,omitempty"`
	OutlierType             CE        `hl7:"11" json:"OutlierType,omitempty"`
	OutlierDays             int       `hl7:"12" json:"OutlierDays,omitempty"`
	OutlierCost             CP        `hl7:"13" json:"OutlierCost,omitempty"`
	GrouperVersionAndType   string    `hl7:"14" json:"GrouperVersionAndType,omitempty"`
	DiagnosisPriority       int       `hl7:"15" json:"DiagnosisPriority,omitempty"`
	DiagnosingClinician     []XCN     `hl7:"16" json:"DiagnosingClinician,omitempty"`
	DiagnosisClassification string    `hl7:"17" json:"DiagnosisClassification,omitempty"`
	ConfidentialIndicator   string    `hl7:"18" json:"ConfidentialIndicator,omitempty"`
	AttestationDateTime     time.Time `hl7:"19,longdate" json:"AttestationDateTime,omitempty"`
}
