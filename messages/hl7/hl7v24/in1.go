package hl7v24

import "time"

// HL7 v2.4 - IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/IN1
type IN1 struct {
	SetID                          string    `hl7:"1" json:"SetID,omitempty"`
	PlanID                         CE        `hl7:"2" json:"PlanID,omitempty"`
	CompanyID                      []CX      `hl7:"3" json:"CompanyID,omitempty"`
	CompanyName                    []XON     `hl7:"4" json:"CompanyName,omitempty"`
	CompanyAddress                 []XAD     `hl7:"5" json:"CompanyAddress,omitempty"`
	ContactPerson                  []XPN     `hl7:"6" json:"ContactPerson,omitempty"`
	PhoneNumber                    []XTN     `hl7:"7" json:"PhoneNumber,omitempty"`
	GroupNumber                    string    `hl7:"8" json:"GroupNumber,omitempty"`
	GroupName                      []XON     `hl7:"9" json:"GroupName,omitempty"`
	GroupEmployerID                []CX      `hl7:"10" json:"GroupEmployerID,omitempty"`
	GroupEmployerName              []XON     `hl7:"11" json:"GroupEmployerName,omitempty"`
	PlanEffectiveDate              time.Time `hl7:"12,shortdate" json:"PlanEffectiveDate,omitempty"`
	PlanExpirationDate             time.Time `hl7:"13,shortdate" json:"PlanExpirationDate,omitempty"`
	AuthorizationInformation       AUI       `hl7:"14" json:"AuthorizationInformation,omitempty"`
	PlanType                       string    `hl7:"15" json:"PlanType,omitempty"`
	NameOfInsured                  []XPN     `hl7:"16" json:"NameOfInsured,omitempty"`
	InsuredRelationshipToPatient   CE        `hl7:"17" json:"InsuredRelationshipToPatient,omitempty"`
	InsuredDateOfBirth             time.Time `hl7:"18,longdate" json:"InsuredDateOfBirth,omitempty"`
	InsuredAddress                 []XAD     `hl7:"19" json:"InsuredAddress,omitempty"`
	AssignmentOfBenefits           string    `hl7:"20" json:"AssignmentOfBenefits,omitempty"`
	CoordinationOfBenefits         string    `hl7:"21" json:"CoordinationOfBenefits,omitempty"`
	CoordOfBenPriority             string    `hl7:"22" json:"CoordOfBenPriority,omitempty"`
	NoticeOfAdmissionFlag          string    `hl7:"23" json:"NoticeOfAdmissionFlag,omitempty"`
	NoticeOfAdmissionDate          time.Time `hl7:"24,shortdate" json:"NoticeOfAdmissionDate,omitempty"`
	ReportOfEligibilityFlag        string    `hl7:"25" json:"ReportOfEligibilityFlag,omitempty"`
	ReportOfEligibilityDate        time.Time `hl7:"26,shortdate" json:"ReportOfEligibilityDate,omitempty"`
	ReleaseInformationCode         string    `hl7:"27" json:"ReleaseInformationCode,omitempty"`
	PreAdmitCert                   string    `hl7:"28" json:"PreAdmitCert,omitempty"`
	VerificationDateTime           time.Time `hl7:"29,longdate" json:"VerificationDateTime,omitempty"`
	VerificationBy                 []XCN     `hl7:"30" json:"VerificationBy,omitempty"`
	TypeOfAgreementCode            string    `hl7:"31" json:"TypeOfAgreementCode,omitempty"`
	BillingStatus                  string    `hl7:"32" json:"BillingStatus,omitempty"`
	LifetimeReserveDays            float32   `hl7:"33" json:"LifetimeReserveDays,omitempty"`
	DelayBeforeLifetimeReserveDays float32   `hl7:"34" json:"DelayBeforeLifetimeReserveDays,omitempty"`
	CompanyPlanCode                string    `hl7:"35" json:"CompanyPlanCode,omitempty"`
	PolicyNumber                   string    `hl7:"36" json:"PolicyNumber,omitempty"`
	PolicyDeductible               CP        `hl7:"37" json:"PolicyDeductible,omitempty"`
	PolicyLimitAmount              CP        `hl7:"38" json:"PolicyLimitAmount,omitempty"`
	PolicyLimitDays                float32   `hl7:"39" json:"PolicyLimitDays,omitempty"`
	RoomRateSemiPrivate            CP        `hl7:"40" json:"RoomRateSemiPrivate,omitempty"`
	RoomRatePrivate                CP        `hl7:"41" json:"RoomRatePrivate,omitempty"`
	InsuredEmploymentStatus        CE        `hl7:"42" json:"InsuredEmploymentStatus,omitempty"`
	InsuredSex                     string    `hl7:"43" json:"InsuredSex,omitempty"`
	InsuredEmployerAddress         []XAD     `hl7:"44" json:"InsuredEmployerAddress,omitempty"`
	VerificationStatus             string    `hl7:"45" json:"VerificationStatus,omitempty"`
	PriorInsurancePlanID           string    `hl7:"46" json:"PriorInsurancePlanID,omitempty"`
	CoverageType                   string    `hl7:"47" json:"CoverageType,omitempty"`
	Handicap                       string    `hl7:"48" json:"Handicap,omitempty"`
	InsuredIDNumber                []CX      `hl7:"49" json:"InsuredIDNumber,omitempty"`
}
