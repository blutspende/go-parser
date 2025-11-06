package hl7v24

import "time"

// HL7 v2.4 - IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/IN1
type IN1 struct {
	SetID                          string    `hl7:"POS=2" json:"SetID,omitempty"`
	PlanID                         CE        `hl7:"POS=3" json:"PlanID,omitempty"`
	CompanyID                      []CX      `hl7:"POS=4" json:"CompanyID,omitempty"`
	CompanyName                    []XON     `hl7:"POS=5" json:"CompanyName,omitempty"`
	CompanyAddress                 []XAD     `hl7:"POS=6" json:"CompanyAddress,omitempty"`
	ContactPerson                  []XPN     `hl7:"POS=7" json:"ContactPerson,omitempty"`
	PhoneNumber                    []XTN     `hl7:"POS=8" json:"PhoneNumber,omitempty"`
	GroupNumber                    string    `hl7:"POS=9" json:"GroupNumber,omitempty"`
	GroupName                      []XON     `hl7:"POS=10" json:"GroupName,omitempty"`
	GroupEmployerID                []CX      `hl7:"POS=11" json:"GroupEmployerID,omitempty"`
	GroupEmployerName              []XON     `hl7:"POS=12" json:"GroupEmployerName,omitempty"`
	PlanEffectiveDate              time.Time `hl7:"POS=13;ATR=date" json:"PlanEffectiveDate,omitempty"`
	PlanExpirationDate             time.Time `hl7:"POS=14;ATR=date" json:"PlanExpirationDate,omitempty"`
	AuthorizationInformation       AUI       `hl7:"POS=15" json:"AuthorizationInformation,omitempty"`
	PlanType                       string    `hl7:"POS=16" json:"PlanType,omitempty"`
	NameOfInsured                  []XPN     `hl7:"POS=17" json:"NameOfInsured,omitempty"`
	InsuredRelationshipToPatient   CE        `hl7:"POS=18" json:"InsuredRelationshipToPatient,omitempty"`
	InsuredDateOfBirth             time.Time `hl7:"POS=19" json:"InsuredDateOfBirth,omitempty"`
	InsuredAddress                 []XAD     `hl7:"POS=20" json:"InsuredAddress,omitempty"`
	AssignmentOfBenefits           string    `hl7:"POS=21" json:"AssignmentOfBenefits,omitempty"`
	CoordinationOfBenefits         string    `hl7:"POS=22" json:"CoordinationOfBenefits,omitempty"`
	CoordOfBenPriority             string    `hl7:"POS=23" json:"CoordOfBenPriority,omitempty"`
	NoticeOfAdmissionFlag          string    `hl7:"POS=24" json:"NoticeOfAdmissionFlag,omitempty"`
	NoticeOfAdmissionDate          time.Time `hl7:"POS=25;ATR=date" json:"NoticeOfAdmissionDate,omitempty"`
	ReportOfEligibilityFlag        string    `hl7:"POS=26" json:"ReportOfEligibilityFlag,omitempty"`
	ReportOfEligibilityDate        time.Time `hl7:"POS=27;ATR=date" json:"ReportOfEligibilityDate,omitempty"`
	ReleaseInformationCode         string    `hl7:"POS=28" json:"ReleaseInformationCode,omitempty"`
	PreAdmitCert                   string    `hl7:"POS=29" json:"PreAdmitCert,omitempty"`
	VerificationDateTime           time.Time `hl7:"POS=30" json:"VerificationDateTime,omitempty"`
	VerificationBy                 []XCN     `hl7:"POS=31" json:"VerificationBy,omitempty"`
	TypeOfAgreementCode            string    `hl7:"POS=32" json:"TypeOfAgreementCode,omitempty"`
	BillingStatus                  string    `hl7:"POS=33" json:"BillingStatus,omitempty"`
	LifetimeReserveDays            *float32  `hl7:"POS=34" json:"LifetimeReserveDays,omitempty"`
	DelayBeforeLifetimeReserveDays *float32  `hl7:"POS=35" json:"DelayBeforeLifetimeReserveDays,omitempty"`
	CompanyPlanCode                string    `hl7:"POS=36" json:"CompanyPlanCode,omitempty"`
	PolicyNumber                   string    `hl7:"POS=37" json:"PolicyNumber,omitempty"`
	PolicyDeductible               CP        `hl7:"POS=38" json:"PolicyDeductible,omitempty"`
	PolicyLimitAmount              CP        `hl7:"POS=39" json:"PolicyLimitAmount,omitempty"`
	PolicyLimitDays                *float32  `hl7:"POS=40" json:"PolicyLimitDays,omitempty"`
	RoomRateSemiPrivate            CP        `hl7:"POS=41" json:"RoomRateSemiPrivate,omitempty"`
	RoomRatePrivate                CP        `hl7:"POS=42" json:"RoomRatePrivate,omitempty"`
	InsuredEmploymentStatus        CE        `hl7:"POS=43" json:"InsuredEmploymentStatus,omitempty"`
	InsuredSex                     string    `hl7:"POS=44" json:"InsuredSex,omitempty"`
	InsuredEmployerAddress         []XAD     `hl7:"POS=45" json:"InsuredEmployerAddress,omitempty"`
	VerificationStatus             string    `hl7:"POS=46" json:"VerificationStatus,omitempty"`
	PriorInsurancePlanID           string    `hl7:"POS=47" json:"PriorInsurancePlanID,omitempty"`
	CoverageType                   string    `hl7:"POS=48" json:"CoverageType,omitempty"`
	Handicap                       string    `hl7:"POS=49" json:"Handicap,omitempty"`
	InsuredIDNumber                []CX      `hl7:"POS=50" json:"InsuredIDNumber,omitempty"`
}
