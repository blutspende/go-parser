package hl7v23

import "time"

// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV2
type PV2 struct {
	PriorPendingLocation              PL        `hl7:"POS=2" json:"priorPendingLocation,omitempty"`
	AccommodationCode                 CE        `hl7:"POS=3" json:"accommodationCode,omitempty"`
	AdmitReason                       CE        `hl7:"POS=4" json:"admitReason,omitempty"`
	TransferReason                    CE        `hl7:"POS=5" json:"transferReason,omitempty"`
	PatientValuables                  string    `hl7:"POS=6" json:"patientValuables,omitempty"`
	PatientValuablesLocation          string    `hl7:"POS=7" json:"patientValuablesLocation,omitempty"`
	VisitUserCode                     string    `hl7:"POS=8" json:"visitUserCode,omitempty"`
	ExpectedAdmitDate                 time.Time `hl7:"POS=9" json:"expectedAdmitDate,omitempty"`
	ExpectedDischargeDate             time.Time `hl7:"POS=10" json:"expectedDischargeDate,omitempty"`
	EstimatedLengthOfInpatientStay    *float32  `hl7:"POS=11" json:"estimatedLengthOfInpatientStay,omitempty"`
	ActualLengthOfImpatientStay       *float32  `hl7:"POS=12" json:"actualLengthOfImpatientStay,omitempty"`
	VisitDescription                  string    `hl7:"POS=13" json:"visitDescription,omitempty"`
	ReferralSourceCode                XCN       `hl7:"POS=14" json:"referralSourceCode,omitempty"`
	PreviousServiceDate               time.Time `hl7:"POS=15" json:"previousServiceDate,omitempty"`
	EmploymentIllnessRelatedIndicator string    `hl7:"POS=16" json:"employmentIllnessRelatedIndicator,omitempty"`
	PurgeStatusCode                   string    `hl7:"POS=17" json:"purgeStatusCode,omitempty"`
	PurgeStatusDate                   time.Time `hl7:"POS=18" json:"purgeStatusDate,omitempty"`
	SpecialProgramCode                string    `hl7:"POS=19" json:"specialProgramCode,omitempty"`
	RetentionIndicator                string    `hl7:"POS=20" json:"retentionIndicator,omitempty"`
	ExpectedNumberOfInsurancePlans    *float32  `hl7:"POS=21" json:"expectedNumberOfInsurancePlans,omitempty"`
	VisitPublicCityCode               string    `hl7:"POS=22" json:"visitPublicCityCode,omitempty"`
	VisitProtectionIndicator          string    `hl7:"POS=23" json:"visitProtectionIndicator,omitempty"`
	ClinicOrganizationName            XON       `hl7:"POS=24" json:"clinicOrganizationName,omitempty"`
	PatientStatusCode                 string    `hl7:"POS=25" json:"patientStatusCode,omitempty"`
	VisitPriorityCode                 string    `hl7:"POS=26" json:"visitPriorityCode,omitempty"`
	PreviousTreatmentCode             string    `hl7:"POS=27" json:"previousTreatmentCode,omitempty"`
	ExpectedDischargeDisposition      string    `hl7:"POS=28" json:"expectedDischargeDisposition,omitempty"`
	SignatureOnFileDate               time.Time `hl7:"POS=29" json:"signatureOnFileDate,omitempty"`
	FirstSimilarIllnessDate           time.Time `hl7:"POS=30" json:"firstSimilarIllnessDate,omitempty"`
	PatientChargeAdjustmentCode       string    `hl7:"POS=31" json:"patientChargeAdjustmentCode,omitempty"`
	RecurringServiceCode              string    `hl7:"POS=32" json:"recurringServiceCode,omitempty"`
	BillingMediaCode                  string    `hl7:"POS=33" json:"billingMediaCode,omitempty"`
	ExpectedSurgeryDateTime           time.Time `hl7:"POS=34" json:"expectedSurgeryDateTime,omitempty"`
	MilitaryPartnershipCode           string    `hl7:"POS=35" json:"militaryPartnershipCode,omitempty"`
	MilitaryNonAvailabilityCode       string    `hl7:"POS=36" json:"militaryNonAvailabilityCode,omitempty"`
	NewbornBabyIndicator              string    `hl7:"POS=37" json:"newbornBabyIndicator,omitempty"`
	BabyDetainedIndicator             string    `hl7:"POS=38" json:"babyDetainedIndicator,omitempty"`
}
