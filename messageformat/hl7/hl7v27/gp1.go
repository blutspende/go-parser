package hl7v27

// GP1 - Grouping/reimbursement - Visit
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/GP1
type GP1 struct {
	TypeOfBillCode CWE `hl7:"POS=2"`
	RevenueCode []CWE `hl7:"POS=3"`
	OverallClaimDispositionCode CWE `hl7:"POS=4"`
	OceEditsPerVisitCode []CWE `hl7:"POS=5"`
	OutlierCost CP `hl7:"POS=6"`
}

