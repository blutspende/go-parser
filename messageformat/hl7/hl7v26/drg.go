package hl7v26

import "time"

// DRG - Diagnosis Related Group
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/DRG
type DRG struct {
	DiagnosticRelatedGroup CNE `hl7:"POS=2"`
	DrgAssignedDateTime time.Time `hl7:"POS=3"`
	DrgApprovalIndicator string `hl7:"POS=4"`
	DrgGrouperReviewCode string `hl7:"POS=5"`
	OutlierType CWE `hl7:"POS=6"`
	OutlierDays *int `hl7:"POS=7"`
	OutlierCost CP `hl7:"POS=8"`
	DrgPayor string `hl7:"POS=9"`
	OutlierReimbursement CP `hl7:"POS=10"`
	ConfidentialIndicator string `hl7:"POS=11"`
	DrgTransferType string `hl7:"POS=12"`
	NameOfCoder XPN `hl7:"POS=13"`
	GrouperStatus CWE `hl7:"POS=14"`
	PcclValueCode CWE `hl7:"POS=15"`
	EffectiveWeight *int `hl7:"POS=16"`
	MonetaryAmount MO `hl7:"POS=17"`
	StatusPatient string `hl7:"POS=18"`
	GrouperSoftwareName string `hl7:"POS=19"`
	GrouperSoftwareVersion string `hl7:"POS=20"`
	StatusFinancialCalculation string `hl7:"POS=21"`
	RelativeDiscountSurcharge MO `hl7:"POS=22"`
	BasicCharge MO `hl7:"POS=23"`
	TotalCharge MO `hl7:"POS=24"`
	DiscountSurcharge MO `hl7:"POS=25"`
	CalculatedDays *int `hl7:"POS=26"`
	StatusGender string `hl7:"POS=27"`
	StatusAge string `hl7:"POS=28"`
	StatusLengthOfStay string `hl7:"POS=29"`
	StatusSameDayFlag string `hl7:"POS=30"`
	StatusSeparationMode string `hl7:"POS=31"`
	StatusWeightAtBirth string `hl7:"POS=32"`
	StatusRespirationMinutes string `hl7:"POS=33"`
	StatusAdmission string `hl7:"POS=34"`
}

