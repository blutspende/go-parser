package hl7v25

// UB2 - UB92 Data
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/UB2
type UB2 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	CoInsuranceDays9 string `hl7:"POS=3"`
	ConditionCode2430 []string `hl7:"POS=4"`
	CoveredDays7 string `hl7:"POS=5"`
	NonCoveredDays8 string `hl7:"POS=6"`
	ValueAmountAndCode []UVC `hl7:"POS=7"`
	OccurrenceCodeAndDate3235 []OCD `hl7:"POS=8"`
	OccurrenceSpanCodeDates36 []OSP `hl7:"POS=9"`
	Ub92Locator2State []string `hl7:"POS=10"`
	Ub92Locator11State []string `hl7:"POS=11"`
	Ub92Locator31National string `hl7:"POS=12"`
	DocumentControlNumber []string `hl7:"POS=13"`
	Ub92Locator49National []string `hl7:"POS=14"`
	Ub92Locator56State []string `hl7:"POS=15"`
	Ub92Locator57National string `hl7:"POS=16"`
	Ub92Locator78State []string `hl7:"POS=17"`
	SpecialVisitCount *int `hl7:"POS=18"`
}

