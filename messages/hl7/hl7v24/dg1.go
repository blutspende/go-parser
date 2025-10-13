package hl7v24

import "time"

// HL7 v2.4 - DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/DG1
type DG1 struct {
	SetID                   int       `hl7:"2,sequence" json:"SetID ,omitempty"`
	DiagnosisCodingMethod   string    `hl7:"3" json:"DiagnosisCodingMethod,omitempty"`
	DiagnosisCode           CE        `hl7:"4" json:"DiagnosisCode,omitempty"`
	DiagnosisDescription    string    `hl7:"5" json:"DiagnosisDescription,omitempty"`
	DiagnosisDateTime       time.Time `hl7:"6,longdate" json:"DiagnosisDateTime,omitempty"`
	DiagnosisType           string    `hl7:"7" json:"DiagnosisType,omitempty"`
	MajorDiagnosticCategory CE        `hl7:"8" json:"MajorDiagnosticCategory,omitempty"`
	DiagnosticRelatedGroup  CE        `hl7:"9" json:"DiagnosticRelatedGroup,omitempty"`
	DRGApprovalIndicator    string    `hl7:"10" json:"DRGApprovalIndicator,omitempty"`
	DRGGrouperReviewCode    string    `hl7:"11" json:"DRGGrouperReviewCode,omitempty"`
	OutlierType             CE        `hl7:"12" json:"OutlierType,omitempty"`
	OutlierDays             int       `hl7:"13" json:"OutlierDays,omitempty"`
	OutlierCost             CP        `hl7:"14" json:"OutlierCost,omitempty"`
	GrouperVersionAndType   string    `hl7:"15" json:"GrouperVersionAndType,omitempty"`
	DiagnosisPriority       int       `hl7:"16" json:"DiagnosisPriority,omitempty"`
	DiagnosingClinician     []XCN     `hl7:"17" json:"DiagnosingClinician,omitempty"`
	DiagnosisClassification string    `hl7:"18" json:"DiagnosisClassification,omitempty"`
	ConfidentialIndicator   string    `hl7:"19" json:"ConfidentialIndicator,omitempty"`
	AttestationDateTime     time.Time `hl7:"20,longdate" json:"AttestationDateTime,omitempty"`
}
