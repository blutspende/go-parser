package hl7v271

// GP2 - Grouping/reimbursement - Procedure Line Item
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/GP2
type GP2 struct {
	RevenueCode CWE `hl7:"POS=2"`
	NumberOfServiceUnits *int `hl7:"POS=3"`
	Charge CP `hl7:"POS=4"`
	ReimbursementActionCode CWE `hl7:"POS=5"`
	DenialOrRejectionCode CWE `hl7:"POS=6"`
	OceEditCode []CWE `hl7:"POS=7"`
	AmbulatoryPaymentClassificationCode CWE `hl7:"POS=8"`
	ModifierEditCode []CWE `hl7:"POS=9"`
	PaymentAdjustmentCode CWE `hl7:"POS=10"`
	PackagingStatusCode CWE `hl7:"POS=11"`
	ExpectedCmsPaymentAmount CP `hl7:"POS=12"`
	ReimbursementTypeCode CWE `hl7:"POS=13"`
	CoPayAmount CP `hl7:"POS=14"`
	PayRatePerServiceUnit *int `hl7:"POS=15"`
}

