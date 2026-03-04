package hl7v28

import "time"

// AUT - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/AUT
type AUT struct {
	AuthorizingPayorPlanID CWE `hl7:"POS=2"`
	AuthorizingPayorCompanyID CWE `hl7:"POS=3"`
	AuthorizingPayorCompanyName string `hl7:"POS=4"`
	AuthorizationEffectiveDate time.Time `hl7:"POS=5"`
	AuthorizationExpirationDate time.Time `hl7:"POS=6"`
	AuthorizationIdentifier EI `hl7:"POS=7"`
	ReimbursementLimit CP `hl7:"POS=8"`
	RequestedNumberOfTreatments CQ `hl7:"POS=9"`
	AuthorizedNumberOfTreatments CQ `hl7:"POS=10"`
	ProcessDate time.Time `hl7:"POS=11"`
	RequestedDisciplines []CWE `hl7:"POS=12"`
	AuthorizedDisciplines []CWE `hl7:"POS=13"`
	AuthorizationReferralType CWE `hl7:"POS=14"`
	ApprovalStatus CWE `hl7:"POS=15"`
	PlannedTreatmentStopDate time.Time `hl7:"POS=16"`
	ClinicalService CWE `hl7:"POS=17"`
	ReasonText string `hl7:"POS=18"`
	NumberOfAuthorizedTreatmentsUnits CQ `hl7:"POS=19"`
	NumberOfUsedTreatmentsUnits CQ `hl7:"POS=20"`
	NumberOfScheduleTreatmentsUnits CQ `hl7:"POS=21"`
	EncounterType CWE `hl7:"POS=22"`
	RemainingBenefitAmount MO `hl7:"POS=23"`
	AuthorizedProvider XON `hl7:"POS=24"`
	AuthorizedHealthProfessional XCN `hl7:"POS=25"`
	SourceText string `hl7:"POS=26"`
	SourceDate time.Time `hl7:"POS=27"`
	SourcePhone XTN `hl7:"POS=28"`
	Comment string `hl7:"POS=29"`
	ActionCode string `hl7:"POS=30"`
}

