package hl7v23

import "time"

// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV2
type PV2 struct {
	PriorPendingLocation              PL        `hl7:"1" json:"priorPendingLocation,omitempty"`
	AccommodationCode                 CE        `hl7:"2" json:"accommodationCode,omitempty"`
	AdmitReason                       CE        `hl7:"3" json:"admitReason,omitempty"`
	TransferReason                    CE        `hl7:"4" json:"transferReason,omitempty"`
	PatientValuables                  string    `hl7:"5" json:"patientValuables,omitempty"`
	PatientValuablesLocation          string    `hl7:"6" json:"patientValuablesLocation,omitempty"`
	VisitUserCode                     string    `hl7:"7" json:"visitUserCode,omitempty"`
	ExpectedAdmitDate                 time.Time `hl7:"8" json:"expectedAdmitDate,omitempty"`
	ExpectedDischargeDate             time.Time `hl7:"9" json:"expectedDischargeDate,omitempty"`
	EstimatedLengthOfInpatientStay    float32   `hl7:"10" json:"estimatedLengthOfInpatientStay,omitempty"`
	ActualLengthOfImpatientStay       float32   `hl7:"11" json:"actualLengthOfImpatientStay,omitempty"`
	VisitDescription                  string    `hl7:"12" json:"visitDescription,omitempty"`
	ReferralSourceCode                XCN       `hl7:"13" json:"referralSourceCode,omitempty"`
	PreviousServiceDate               time.Time `hl7:"14" json:"previousServiceDate,omitempty"`
	EmploymentIllnessRelatedIndicator string    `hl7:"15" json:"employmentIllnessRelatedIndicator,omitempty"`
	PurgeStatusCode                   string    `hl7:"16" json:"purgeStatusCode,omitempty"`
	PurgeStatusDate                   time.Time `hl7:"17" json:"purgeStatusDate,omitempty"`
	SpecialProgramCode                string    `hl7:"18" json:"specialProgramCode,omitempty"`
	RetentionIndicator                string    `hl7:"19" json:"retentionIndicator,omitempty"`
	ExpectedNumberOfInsurancePlans    float32   `hl7:"20" json:"expectedNumberOfInsurancePlans,omitempty"`
	VisitPublicCityCode               string    `hl7:"21" json:"visitPublicCityCode,omitempty"`
	VisitProtectionIndicator          string    `hl7:"22" json:"visitProtectionIndicator,omitempty"`
	ClinicOrganizationName            XON       `hl7:"23" json:"clinicOrganizationName,omitempty"`
	PatientStatusCode                 string    `hl7:"24" json:"patientStatusCode,omitempty"`
	VisitPriorityCode                 string    `hl7:"25" json:"visitPriorityCode,omitempty"`
	PreviousTreatmentCode             string    `hl7:"26" json:"previousTreatmentCode,omitempty"`
	ExpectedDischargeDisposition      string    `hl7:"27" json:"expectedDischargeDisposition,omitempty"`
	SignatureOnFileDate               time.Time `hl7:"28" json:"signatureOnFileDate,omitempty"`
	FirstSimilarIllnessDate           time.Time `hl7:"29" json:"firstSimilarIllnessDate,omitempty"`
	PatientChargeAdjustmentCode       string    `hl7:"30" json:"patientChargeAdjustmentCode,omitempty"`
	RecurringServiceCode              string    `hl7:"31" json:"recurringServiceCode,omitempty"`
	BillingMediaCode                  string    `hl7:"32" json:"billingMediaCode,omitempty"`
	ExpectedSurgeryDateTime           time.Time `hl7:"33" json:"expectedSurgeryDateTime,omitempty"`
	MilitaryPartnershipCode           string    `hl7:"34" json:"militaryPartnershipCode,omitempty"`
	MilitaryNonAvailabilityCode       string    `hl7:"35" json:"militaryNonAvailabilityCode,omitempty"`
	NewbornBabyIndicator              string    `hl7:"36" json:"newbornBabyIndicator,omitempty"`
	BabyDetainedIndicator             string    `hl7:"37" json:"babyDetainedIndicator,omitempty"`
}
