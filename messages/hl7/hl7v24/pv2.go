package hl7v24

import "time"

// HL7 v2.4 - PV2 - Patient visit - additional information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PV2
type PV2 struct {
	PriorPendingLocation              PL        `hl7:"2" json:"PriorPendingLocation,omitempty"`
	AccommodationCode                 CE        `hl7:"3" json:"AccommodationCode,omitempty"`
	AdmitReason                       CE        `hl7:"4" json:"AdmitReason,omitempty"`
	TransferReason                    CE        `hl7:"5" json:"TransferReason,omitempty"`
	PatientValuables                  string    `hl7:"6" json:"PatientValuables,omitempty"`
	PatientValuablesLocation          string    `hl7:"7" json:"PatientValuablesLocation,omitempty"`
	VisitUserCode                     string    `hl7:"8" json:"VisitUserCode,omitempty"`
	ExpectedAdmitDate                 time.Time `hl7:"9" json:"ExpectedAdmitDate,omitempty"`
	ExpectedDischargeDate             time.Time `hl7:"10" json:"ExpectedDischargeDate,omitempty"`
	EstimatedLengthOfInpatientStay    float32   `hl7:"11" json:"EstimatedLengthOfInpatientStay,omitempty"`
	ActualLengthOfImpatientStay       float32   `hl7:"12" json:"ActualLengthOfImpatientStay,omitempty"`
	VisitDescription                  string    `hl7:"13" json:"VisitDescription,omitempty"`
	ReferralSourceCode                []XCN     `hl7:"14" json:"ReferralSourceCode,omitempty"`
	PreviousServiceDate               time.Time `hl7:"15" json:"PreviousServiceDate,omitempty"`
	EmploymentIllnessRelatedIndicator string    `hl7:"16" json:"EmploymentIllnessRelatedIndicator,omitempty"`
	PurgeStatusCode                   string    `hl7:"17" json:"PurgeStatusCode,omitempty"`
	PurgeStatusDate                   time.Time `hl7:"18" json:"PurgeStatusDate,omitempty"`
	SpecialProgramCode                string    `hl7:"19" json:"SpecialProgramCode,omitempty"`
	RetentionIndicator                string    `hl7:"20" json:"RetentionIndicator,omitempty"`
	ExpectedNumberOfInsurancePlans    float32   `hl7:"21" json:"ExpectedNumberOfInsurancePlans,omitempty"`
	VisitPublicityCode                string    `hl7:"22" json:"VisitPublicityCode,omitempty"`
	VisitProtectionIndicator          string    `hl7:"23" json:"VisitProtectionIndicator,omitempty"`
	ClinicOrganizationName            XON       `hl7:"24" json:"ClinicOrganizationName,omitempty"`
	PatientStatusCode                 string    `hl7:"25" json:"PatientStatusCode,omitempty"`
	VisitPriorityCode                 string    `hl7:"26" json:"VisitPriorityCode,omitempty"`
	PreviousTreatmentCode             string    `hl7:"27" json:"PreviousTreatmentCode,omitempty"`
	ExpectedDischargeDisposition      string    `hl7:"28" json:"ExpectedDischargeDisposition,omitempty"`
	SignatureOnFileDate               time.Time `hl7:"29" json:"SignatureOnFileDate,omitempty"`
	FirstSimilarIllnessDate           time.Time `hl7:"30" json:"FirstSimilarIllnessDate,omitempty"`
	PatientChargeAdjustmentCode       CE        `hl7:"31" json:"PatientChargeAdjustmentCode,omitempty"`
	RecurringServiceCode              string    `hl7:"32" json:"RecurringServiceCode,omitempty"`
	BillingMediaCode                  string    `hl7:"33" json:"BillingMediaCode,omitempty"`
	ExpectedSurgeryDateTime           time.Time `hl7:"34" json:"ExpectedSurgeryDateTime,omitempty"`
	MilitaryPartnershipCode           string    `hl7:"35" json:"MilitaryPartnershipCode,omitempty"`
	MilitaryNonAvailabilityCode       string    `hl7:"36" json:"MilitaryNonAvailabilityCode,omitempty"`
	NewbornBabyIndicator              string    `hl7:"37" json:"NewbornBabyIndicator,omitempty"`
	BabyDetainedIndicator             string    `hl7:"38" json:"BabyDetainedIndicator,omitempty"`
	ModeOfArrivalCode                 CE        `hl7:"39" json:"ModeOfArrivalCode,omitempty"`
	RecreationalDrugUseCode           []CE      `hl7:"40" json:"RecreationalDrugUseCode,omitempty"`
	AdmissionLevelOfCareCode          CE        `hl7:"41" json:"AdmissionLevelOfCareCode,omitempty"`
	PrecautionCode                    []CE      `hl7:"42" json:"PrecautionCode,omitempty"`
	PatientConditionCode              CE        `hl7:"43" json:"PatientConditionCode,omitempty"`
	LivingWillCode                    string    `hl7:"44" json:"LivingWillCode,omitempty"`
	OrganDonorCode                    string    `hl7:"45" json:"OrganDonorCode,omitempty"`
	AdvanceDirectiveCode              []CE      `hl7:"46" json:"AdvanceDirectiveCode,omitempty"`
	PatientStatusEffectiveDate        time.Time `hl7:"47" json:"PatientStatusEffectiveDate,omitempty"`
	ExpectedLOAReturnDateTime         time.Time `hl7:"48,longdate" json:"ExpectedLOAReturnDateTime,omitempty"`
}
