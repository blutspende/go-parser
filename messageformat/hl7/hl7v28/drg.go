package hl7v28

import "time"

// DRG - Diagnosis Related Group
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/DRG
type DRG struct {
	DiagnosticRelatedGroup CNE `hl7:"POS=2"`
	DrgAssignedDateTime time.Time `hl7:"POS=3"`
	DrgApprovalIndicator string `hl7:"POS=4"`
	DrgGrouperReviewCode CWE `hl7:"POS=5"`
	OutlierType CWE `hl7:"POS=6"`
	OutlierDays *int `hl7:"POS=7"`
	OutlierCost CP `hl7:"POS=8"`
	DrgPayor CWE `hl7:"POS=9"`
	OutlierReimbursement CP `hl7:"POS=10"`
	ConfidentialIndicator string `hl7:"POS=11"`
	DrgTransferType CWE `hl7:"POS=12"`
	NameOfCoder XPN `hl7:"POS=13"`
	GrouperStatus CWE `hl7:"POS=14"`
	PcclValueCode CWE `hl7:"POS=15"`
	EffectiveWeight *int `hl7:"POS=16"`
	MonetaryAmount MO `hl7:"POS=17"`
	StatusPatient CWE `hl7:"POS=18"`
	GrouperSoftwareName string `hl7:"POS=19"`
	GrouperSoftwareVersion string `hl7:"POS=20"`
	StatusFinancialCalculation CWE `hl7:"POS=21"`
	RelativeDiscountSurcharge MO `hl7:"POS=22"`
	BasicCharge MO `hl7:"POS=23"`
	TotalCharge MO `hl7:"POS=24"`
	DiscountSurcharge MO `hl7:"POS=25"`
	CalculatedDays *int `hl7:"POS=26"`
	StatusGender CWE `hl7:"POS=27"`
	StatusAge CWE `hl7:"POS=28"`
	StatusLengthOfStay CWE `hl7:"POS=29"`
	StatusSameDayFlag CWE `hl7:"POS=30"`
	StatusSeparationMode CWE `hl7:"POS=31"`
	StatusWeightAtBirth CWE `hl7:"POS=32"`
	StatusRespirationMinutes CWE `hl7:"POS=33"`
	StatusAdmission CWE `hl7:"POS=34"`
}

