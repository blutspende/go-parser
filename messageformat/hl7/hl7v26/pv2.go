package hl7v26

import "time"

// PV2 - Patient Visit - Additional Information
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/PV2
type PV2 struct {
	PriorPendingLocation PL `hl7:"POS=2"`
	AccommodationCode CWE `hl7:"POS=3"`
	AdmitReason CWE `hl7:"POS=4"`
	TransferReason CWE `hl7:"POS=5"`
	PatientValuables []string `hl7:"POS=6"`
	PatientValuablesLocation string `hl7:"POS=7"`
	VisitUserCode []string `hl7:"POS=8"`
	ExpectedAdmitDateTime time.Time `hl7:"POS=9"`
	ExpectedDischargeDateTime time.Time `hl7:"POS=10"`
	EstimatedLengthOfInpatientStay *int `hl7:"POS=11"`
	ActualLengthOfInpatientStay *int `hl7:"POS=12"`
	VisitDescription string `hl7:"POS=13"`
	ReferralSourceCode []XCN `hl7:"POS=14"`
	PreviousServiceDate time.Time `hl7:"POS=15;ATR=date"`
	EmploymentIllnessRelatedIndicator string `hl7:"POS=16"`
	PurgeStatusCode string `hl7:"POS=17"`
	PurgeStatusDate time.Time `hl7:"POS=18;ATR=date"`
	SpecialProgramCode string `hl7:"POS=19"`
	RetentionIndicator string `hl7:"POS=20"`
	ExpectedNumberOfInsurancePlans *int `hl7:"POS=21"`
	VisitPublicityCode string `hl7:"POS=22"`
	VisitProtectionIndicator string `hl7:"POS=23"`
	ClinicOrganizationName []XON `hl7:"POS=24"`
	PatientStatusCode string `hl7:"POS=25"`
	VisitPriorityCode string `hl7:"POS=26"`
	PreviousTreatmentDate time.Time `hl7:"POS=27;ATR=date"`
	ExpectedDischargeDisposition string `hl7:"POS=28"`
	SignatureOnFileDate time.Time `hl7:"POS=29;ATR=date"`
	FirstSimilarIllnessDate time.Time `hl7:"POS=30;ATR=date"`
	PatientChargeAdjustmentCode CWE `hl7:"POS=31"`
	RecurringServiceCode string `hl7:"POS=32"`
	BillingMediaCode string `hl7:"POS=33"`
	ExpectedSurgeryDateAndTime time.Time `hl7:"POS=34"`
	MilitaryPartnershipCode string `hl7:"POS=35"`
	MilitaryNonAvailabilityCode string `hl7:"POS=36"`
	NewbornBabyIndicator string `hl7:"POS=37"`
	BabyDetainedIndicator string `hl7:"POS=38"`
	ModeOfArrivalCode CWE `hl7:"POS=39"`
	RecreationalDrugUseCode []CWE `hl7:"POS=40"`
	AdmissionLevelOfCareCode CWE `hl7:"POS=41"`
	PrecautionCode []CWE `hl7:"POS=42"`
	PatientConditionCode CWE `hl7:"POS=43"`
	LivingWillCode string `hl7:"POS=44"`
	OrganDonorCode string `hl7:"POS=45"`
	AdvanceDirectiveCode []CWE `hl7:"POS=46"`
	PatientStatusEffectiveDate time.Time `hl7:"POS=47;ATR=date"`
	ExpectedLoaReturnDateTime time.Time `hl7:"POS=48"`
	ExpectedPreAdmissionTestingDateTime time.Time `hl7:"POS=49"`
	NotifyClergyCode []string `hl7:"POS=50"`
	AdvanceDirectiveLastVerifiedDate time.Time `hl7:"POS=51;ATR=date"`
}

