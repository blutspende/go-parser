package hl7v23

import "time"

// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV2
type PV2 struct {
	PriorPendingLocation              PL        `hl7:"2" json:"priorPendingLocation,omitempty"`
	AccommodationCode                 CE        `hl7:"3" json:"accommodationCode,omitempty"`
	AdmitReason                       CE        `hl7:"4" json:"admitReason,omitempty"`
	TransferReason                    CE        `hl7:"5" json:"transferReason,omitempty"`
	PatientValuables                  string    `hl7:"6" json:"patientValuables,omitempty"`
	PatientValuablesLocation          string    `hl7:"7" json:"patientValuablesLocation,omitempty"`
	VisitUserCode                     string    `hl7:"8" json:"visitUserCode,omitempty"`
	ExpectedAdmitDate                 time.Time `hl7:"9" json:"expectedAdmitDate,omitempty"`
	ExpectedDischargeDate             time.Time `hl7:"10"json:"expectedDischargeDate,omitempty"`
	EstimatedLengthOfInpatientStay    float32   `hl7:"11" json:"estimatedLengthOfInpatientStay,omitempty"`
	ActualLengthOfImpatientStay       float32   `hl7:"12" json:"actualLengthOfImpatientStay,omitempty"`
	VisitDescription                  string    `hl7:"13" json:"visitDescription,omitempty"`
	ReferralSourceCode                XCN       `hl7:"14" json:"referralSourceCode,omitempty"`
	PreviousServiceDate               time.Time `hl7:"15" json:"previousServiceDate,omitempty"`
	EmploymentIllnessRelatedIndicator string    `hl7:"16" json:"employmentIllnessRelatedIndicator,omitempty"`
	PurgeStatusCode                   string    `hl7:"17" json:"purgeStatusCode,omitempty"`
	PurgeStatusDate                   time.Time `hl7:"18" json:"purgeStatusDate,omitempty"`
	SpecialProgramCode                string    `hl7:"19" json:"specialProgramCode,omitempty"`
	RetentionIndicator                string    `hl7:"20" json:"retentionIndicator,omitempty"`
	ExpectedNumberOfInsurancePlans    float32   `hl7:"21" json:"expectedNumberOfInsurancePlans,omitempty"`
	VisitPublicCityCode               string    `hl7:"22" json:"visitPublicCityCode,omitempty"`
	VisitProtectionIndicator          string    `hl7:"23" json:"visitProtectionIndicator,omitempty"`
	ClinicOrganizationName            XON       `hl7:"24" json:"clinicOrganizationName,omitempty"`
	PatientStatusCode                 string    `hl7:"25" json:"patientStatusCode,omitempty"`
	VisitPriorityCode                 string    `hl7:"26" json:"visitPriorityCode,omitempty"`
	PreviousTreatmentCode             string    `hl7:"27" json:"previousTreatmentCode,omitempty"`
	ExpectedDischargeDisposition      string    `hl7:"28" json:"expectedDischargeDisposition,omitempty"`
	SignatureOnFileDate               time.Time `hl7:"29" json:"signatureOnFileDate,omitempty"`
	FirstSimilarIllnessDate           time.Time `hl7:"30" json:"firstSimilarIllnessDate,omitempty"`
	PatientChargeAdjustmentCode       string    `hl7:"31" json:"patientChargeAdjustmentCode,omitempty"`
	RecurringServiceCode              string    `hl7:"32" json:"recurringServiceCode,omitempty"`
	BillingMediaCode                  string    `hl7:"33" json:"billingMediaCode,omitempty"`
	ExpectedSurgeryDateTime           time.Time `hl7:"34" json:"expectedSurgeryDateTime,omitempty"`
	MilitaryPartnershipCode           string    `hl7:"35" json:"militaryPartnershipCode,omitempty"`
	MilitaryNonAvailabilityCode       string    `hl7:"36" json:"militaryNonAvailabilityCode,omitempty"`
	NewbornBabyIndicator              string    `hl7:"37" json:"newbornBabyIndicator,omitempty"`
	BabyDetainedIndicator             string    `hl7:"38" json:"babyDetainedIndicator,omitempty"`
}
