package hl7v24

import "time"

// HL7 v2.4 - PV2 - Patient visit - additional information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PV2
type PV2 struct {
	PriorPendingLocation              PL        `hl7:"POS=2" json:"PriorPendingLocation,omitempty"`
	AccommodationCode                 CE        `hl7:"POS=3" json:"AccommodationCode,omitempty"`
	AdmitReason                       CE        `hl7:"POS=4" json:"AdmitReason,omitempty"`
	TransferReason                    CE        `hl7:"POS=5" json:"TransferReason,omitempty"`
	PatientValuables                  string    `hl7:"POS=6" json:"PatientValuables,omitempty"`
	PatientValuablesLocation          string    `hl7:"POS=7" json:"PatientValuablesLocation,omitempty"`
	VisitUserCode                     string    `hl7:"POS=8" json:"VisitUserCode,omitempty"`
	ExpectedAdmitDate                 time.Time `hl7:"POS=9" json:"ExpectedAdmitDate,omitempty"`
	ExpectedDischargeDate             time.Time `hl7:"POS=10" json:"ExpectedDischargeDate,omitempty"`
	EstimatedLengthOfInpatientStay    *float32  `hl7:"POS=11" json:"EstimatedLengthOfInpatientStay,omitempty"`
	ActualLengthOfImpatientStay       *float32  `hl7:"POS=12" json:"ActualLengthOfImpatientStay,omitempty"`
	VisitDescription                  string    `hl7:"POS=13" json:"VisitDescription,omitempty"`
	ReferralSourceCode                []XCN     `hl7:"POS=14" json:"ReferralSourceCode,omitempty"`
	PreviousServiceDate               time.Time `hl7:"POS=15" json:"PreviousServiceDate,omitempty"`
	EmploymentIllnessRelatedIndicator string    `hl7:"POS=16" json:"EmploymentIllnessRelatedIndicator,omitempty"`
	PurgeStatusCode                   string    `hl7:"POS=17" json:"PurgeStatusCode,omitempty"`
	PurgeStatusDate                   time.Time `hl7:"POS=18" json:"PurgeStatusDate,omitempty"`
	SpecialProgramCode                string    `hl7:"POS=19" json:"SpecialProgramCode,omitempty"`
	RetentionIndicator                string    `hl7:"POS=20" json:"RetentionIndicator,omitempty"`
	ExpectedNumberOfInsurancePlans    *float32  `hl7:"POS=21" json:"ExpectedNumberOfInsurancePlans,omitempty"`
	VisitPublicityCode                string    `hl7:"POS=22" json:"VisitPublicityCode,omitempty"`
	VisitProtectionIndicator          string    `hl7:"POS=23" json:"VisitProtectionIndicator,omitempty"`
	ClinicOrganizationName            XON       `hl7:"POS=24" json:"ClinicOrganizationName,omitempty"`
	PatientStatusCode                 string    `hl7:"POS=25" json:"PatientStatusCode,omitempty"`
	VisitPriorityCode                 string    `hl7:"POS=26" json:"VisitPriorityCode,omitempty"`
	PreviousTreatmentCode             string    `hl7:"POS=27" json:"PreviousTreatmentCode,omitempty"`
	ExpectedDischargeDisposition      string    `hl7:"POS=28" json:"ExpectedDischargeDisposition,omitempty"`
	SignatureOnFileDate               time.Time `hl7:"POS=29" json:"SignatureOnFileDate,omitempty"`
	FirstSimilarIllnessDate           time.Time `hl7:"POS=30" json:"FirstSimilarIllnessDate,omitempty"`
	PatientChargeAdjustmentCode       CE        `hl7:"POS=31" json:"PatientChargeAdjustmentCode,omitempty"`
	RecurringServiceCode              string    `hl7:"POS=32" json:"RecurringServiceCode,omitempty"`
	BillingMediaCode                  string    `hl7:"POS=33" json:"BillingMediaCode,omitempty"`
	ExpectedSurgeryDateTime           time.Time `hl7:"POS=34" json:"ExpectedSurgeryDateTime,omitempty"`
	MilitaryPartnershipCode           string    `hl7:"POS=35" json:"MilitaryPartnershipCode,omitempty"`
	MilitaryNonAvailabilityCode       string    `hl7:"POS=36" json:"MilitaryNonAvailabilityCode,omitempty"`
	NewbornBabyIndicator              string    `hl7:"POS=37" json:"NewbornBabyIndicator,omitempty"`
	BabyDetainedIndicator             string    `hl7:"POS=38" json:"BabyDetainedIndicator,omitempty"`
	ModeOfArrivalCode                 CE        `hl7:"POS=39" json:"ModeOfArrivalCode,omitempty"`
	RecreationalDrugUseCode           []CE      `hl7:"POS=40" json:"RecreationalDrugUseCode,omitempty"`
	AdmissionLevelOfCareCode          CE        `hl7:"POS=41" json:"AdmissionLevelOfCareCode,omitempty"`
	PrecautionCode                    []CE      `hl7:"POS=42" json:"PrecautionCode,omitempty"`
	PatientConditionCode              CE        `hl7:"POS=43" json:"PatientConditionCode,omitempty"`
	LivingWillCode                    string    `hl7:"POS=44" json:"LivingWillCode,omitempty"`
	OrganDonorCode                    string    `hl7:"POS=45" json:"OrganDonorCode,omitempty"`
	AdvanceDirectiveCode              []CE      `hl7:"POS=46" json:"AdvanceDirectiveCode,omitempty"`
	PatientStatusEffectiveDate        time.Time `hl7:"POS=47;ATR=date" json:"PatientStatusEffectiveDate,omitempty"`
	ExpectedLOAReturnDateTime         time.Time `hl7:"POS=48" json:"ExpectedLOAReturnDateTime,omitempty"`
}
