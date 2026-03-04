package hl7v24

// GP1 - Grouping/Reimbursement - Visit
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/GP1
type GP1 struct {
	TypeOfBillCode string `hl7:"POS=2"`
	RevenueCode []string `hl7:"POS=3"`
	OverallClaimDispositionCode string `hl7:"POS=4"`
	OceEditsPerVisitCode []string `hl7:"POS=5"`
	OutlierCost CP `hl7:"POS=6"`
}

