package hl7v231

// UB2 - UB92 data segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/UB2
type UB2 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	CoInsuranceDays string `hl7:"POS=3"`
	ConditionCode []string `hl7:"POS=4"`
	CoveredDays string `hl7:"POS=5"`
	NonCoveredDays string `hl7:"POS=6"`
	ValueAmountCode []UVC `hl7:"POS=7"`
	OccurrenceCodeDate []OCD `hl7:"POS=8"`
	OccurrenceSpanCodeDates []OSP `hl7:"POS=9"`
	Ub92Locator2 []string `hl7:"POS=10"`
	Ub92Locator11 []string `hl7:"POS=11"`
	Ub92Locator31 string `hl7:"POS=12"`
	DocumentControlNumber []string `hl7:"POS=13"`
	Ub92Locator49 []string `hl7:"POS=14"`
	Ub92Locator56 []string `hl7:"POS=15"`
	Ub92Locator57 string `hl7:"POS=16"`
	Ub92Locator78 []string `hl7:"POS=17"`
	SpecialVisitCount *int `hl7:"POS=18"`
}

