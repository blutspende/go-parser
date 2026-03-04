package hl7v25

import "time"

// UB1 - UB82 Data
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/UB1
type UB1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	BloodDeductible43 *int `hl7:"POS=3"`
	BloodFurnishedPintsOf40 *int `hl7:"POS=4"`
	BloodReplacedPints41 *int `hl7:"POS=5"`
	BloodNotReplacedPints42 *int `hl7:"POS=6"`
	CoInsuranceDays25 *int `hl7:"POS=7"`
	ConditionCode3539 []string `hl7:"POS=8"`
	CoveredDays23 *int `hl7:"POS=9"`
	NonCoveredDays24 *int `hl7:"POS=10"`
	ValueAmountAndCode4649 []UVC `hl7:"POS=11"`
	NumberOfGraceDays90 *int `hl7:"POS=12"`
	SpecialProgramIndicator44 CE `hl7:"POS=13"`
	PsroUrApprovalIndicator87 CE `hl7:"POS=14"`
	PsroUrApprovedStayFm88 time.Time `hl7:"POS=15;ATR=date"`
	PsroUrApprovedStayTo89 time.Time `hl7:"POS=16;ATR=date"`
	Occurrence2832 []OCD `hl7:"POS=17"`
	OccurrenceSpan33 CE `hl7:"POS=18"`
	OccurSpanStartDate33 time.Time `hl7:"POS=19;ATR=date"`
	OccurSpanEndDate33 time.Time `hl7:"POS=20;ATR=date"`
	Ub82Locator2 string `hl7:"POS=21"`
	Ub82Locator9 string `hl7:"POS=22"`
	Ub82Locator27 string `hl7:"POS=23"`
	Ub82Locator45 string `hl7:"POS=24"`
}

