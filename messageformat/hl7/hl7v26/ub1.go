package hl7v26

import "time"

// UB1 - Uniform Billing 1
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/UB1
type UB1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	BloodFurnishedPints *int `hl7:"POS=4"`
	BloodReplacedPints *int `hl7:"POS=5"`
	BloodNotReplacedPints *int `hl7:"POS=6"`
	CoInsuranceDays *int `hl7:"POS=7"`
	ConditionCode []string `hl7:"POS=8"`
	CoveredDays *int `hl7:"POS=9"`
	NonCoveredDays *int `hl7:"POS=10"`
	ValueAmountCode []UVC `hl7:"POS=11"`
	NumberOfGraceDays *int `hl7:"POS=12"`
	SpecialProgramIndicator CWE `hl7:"POS=13"`
	PsroUrApprovalIndicator CWE `hl7:"POS=14"`
	PsroUrApprovedStayFm time.Time `hl7:"POS=15;ATR=date"`
	PsroUrApprovedStayTo time.Time `hl7:"POS=16;ATR=date"`
	Occurrence []OCD `hl7:"POS=17"`
	OccurrenceSpan CWE `hl7:"POS=18"`
	OccurSpanStartDate time.Time `hl7:"POS=19;ATR=date"`
	OccurSpanEndDate time.Time `hl7:"POS=20;ATR=date"`
	Ub82Locator2 string `hl7:"POS=21"`
	Ub82Locator9 string `hl7:"POS=22"`
	Ub82Locator27 string `hl7:"POS=23"`
	Ub82Locator45 string `hl7:"POS=24"`
}

