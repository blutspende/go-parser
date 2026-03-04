package hl7v24

import "time"

// DRG - Diagnosis Related Group
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/DRG
type DRG struct {
	DiagnosticRelatedGroup CE `hl7:"POS=2"`
	DrgAssignedDateTime time.Time `hl7:"POS=3"`
	DrgApprovalIndicator string `hl7:"POS=4"`
	DrgGrouperReviewCode string `hl7:"POS=5"`
	OutlierType CE `hl7:"POS=6"`
	OutlierDays *int `hl7:"POS=7"`
	OutlierCost CP `hl7:"POS=8"`
	DrgPayor string `hl7:"POS=9"`
	OutlierReimbursement CP `hl7:"POS=10"`
	ConfidentialIndicator string `hl7:"POS=11"`
	DrgTransferType string `hl7:"POS=12"`
}

