package hl7v24

import "time"

// HL7 v2.4 - IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/IN1
type IN1 struct {
	SetID                          string    `hl7:"2" json:"SetID,omitempty"`
	PlanID                         CE        `hl7:"3" json:"PlanID,omitempty"`
	CompanyID                      []CX      `hl7:"4" json:"CompanyID,omitempty"`
	CompanyName                    []XON     `hl7:"5" json:"CompanyName,omitempty"`
	CompanyAddress                 []XAD     `hl7:"6" json:"CompanyAddress,omitempty"`
	ContactPerson                  []XPN     `hl7:"7" json:"ContactPerson,omitempty"`
	PhoneNumber                    []XTN     `hl7:"8" json:"PhoneNumber,omitempty"`
	GroupNumber                    string    `hl7:"9" json:"GroupNumber,omitempty"`
	GroupName                      []XON     `hl7:"10" json:"GroupName,omitempty"`
	GroupEmployerID                []CX      `hl7:"11" json:"GroupEmployerID,omitempty"`
	GroupEmployerName              []XON     `hl7:"12" json:"GroupEmployerName,omitempty"`
	PlanEffectiveDate              time.Time `hl7:"13" json:"PlanEffectiveDate,omitempty"`
	PlanExpirationDate             time.Time `hl7:"14" json:"PlanExpirationDate,omitempty"`
	AuthorizationInformation       AUI       `hl7:"15" json:"AuthorizationInformation,omitempty"`
	PlanType                       string    `hl7:"16" json:"PlanType,omitempty"`
	NameOfInsured                  []XPN     `hl7:"17" json:"NameOfInsured,omitempty"`
	InsuredRelationshipToPatient   CE        `hl7:"18" json:"InsuredRelationshipToPatient,omitempty"`
	InsuredDateOfBirth             time.Time `hl7:"19,longdate" json:"InsuredDateOfBirth,omitempty"`
	InsuredAddress                 []XAD     `hl7:"20" json:"InsuredAddress,omitempty"`
	AssignmentOfBenefits           string    `hl7:"21" json:"AssignmentOfBenefits,omitempty"`
	CoordinationOfBenefits         string    `hl7:"22" json:"CoordinationOfBenefits,omitempty"`
	CoordOfBenPriority             string    `hl7:"23" json:"CoordOfBenPriority,omitempty"`
	NoticeOfAdmissionFlag          string    `hl7:"24" json:"NoticeOfAdmissionFlag,omitempty"`
	NoticeOfAdmissionDate          time.Time `hl7:"25" json:"NoticeOfAdmissionDate,omitempty"`
	ReportOfEligibilityFlag        string    `hl7:"26" json:"ReportOfEligibilityFlag,omitempty"`
	ReportOfEligibilityDate        time.Time `hl7:"27" json:"ReportOfEligibilityDate,omitempty"`
	ReleaseInformationCode         string    `hl7:"28" json:"ReleaseInformationCode,omitempty"`
	PreAdmitCert                   string    `hl7:"29" json:"PreAdmitCert,omitempty"`
	VerificationDateTime           time.Time `hl7:"30,longdate" json:"VerificationDateTime,omitempty"`
	VerificationBy                 []XCN     `hl7:"31" json:"VerificationBy,omitempty"`
	TypeOfAgreementCode            string    `hl7:"32" json:"TypeOfAgreementCode,omitempty"`
	BillingStatus                  string    `hl7:"33" json:"BillingStatus,omitempty"`
	LifetimeReserveDays            float32   `hl7:"34" json:"LifetimeReserveDays,omitempty"`
	DelayBeforeLifetimeReserveDays float32   `hl7:"35" json:"DelayBeforeLifetimeReserveDays,omitempty"`
	CompanyPlanCode                string    `hl7:"36" json:"CompanyPlanCode,omitempty"`
	PolicyNumber                   string    `hl7:"37" json:"PolicyNumber,omitempty"`
	PolicyDeductible               CP        `hl7:"38" json:"PolicyDeductible,omitempty"`
	PolicyLimitAmount              CP        `hl7:"39" json:"PolicyLimitAmount,omitempty"`
	PolicyLimitDays                float32   `hl7:"40" json:"PolicyLimitDays,omitempty"`
	RoomRateSemiPrivate            CP        `hl7:"41" json:"RoomRateSemiPrivate,omitempty"`
	RoomRatePrivate                CP        `hl7:"42" json:"RoomRatePrivate,omitempty"`
	InsuredEmploymentStatus        CE        `hl7:"43" json:"InsuredEmploymentStatus,omitempty"`
	InsuredSex                     string    `hl7:"44" json:"InsuredSex,omitempty"`
	InsuredEmployerAddress         []XAD     `hl7:"45" json:"InsuredEmployerAddress,omitempty"`
	VerificationStatus             string    `hl7:"46" json:"VerificationStatus,omitempty"`
	PriorInsurancePlanID           string    `hl7:"47" json:"PriorInsurancePlanID,omitempty"`
	CoverageType                   string    `hl7:"48" json:"CoverageType,omitempty"`
	Handicap                       string    `hl7:"49" json:"Handicap,omitempty"`
	InsuredIDNumber                []CX      `hl7:"50" json:"InsuredIDNumber,omitempty"`
}
