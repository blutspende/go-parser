package hl7v22

import "time"

// IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/IN1
type IN1 struct {
	SetID                          int       `hl7:"POS=2;ATR=sequence"`
	InsurancePlanID                string    `hl7:"POS=3"`
	InsuranceCompanyID             string    `hl7:"POS=4"`
	InsuranceCompanyName           string    `hl7:"POS=5"`
	InsuranceCompanyAddress        AD        `hl7:"POS=6"`
	InsuranceCompanyContactPers    PN        `hl7:"POS=7"`
	InsuranceCompanyPhoneNumber    []string  `hl7:"POS=8"`
	GroupNumber                    string    `hl7:"POS=9"`
	GroupName                      string    `hl7:"POS=10"`
	InsuredsGroupEmployerID        string    `hl7:"POS=11"`
	InsuredsGroupEmployerName      string    `hl7:"POS=12"`
	PlanEffectiveDate              time.Time `hl7:"POS=13;ATR=date"`
	PlanExpirationDate             time.Time `hl7:"POS=14;ATR=date"`
	AuthorizationInformation       CM_AUI    `hl7:"POS=15"`
	PlanType                       string    `hl7:"POS=16"`
	NameOfInsured                  PN        `hl7:"POS=17"`
	InsuredsRelationshipToPatient  string    `hl7:"POS=18"`
	InsuredsDateOfBirth            time.Time `hl7:"POS=19;ATR=date"`
	InsuredsAddress                AD        `hl7:"POS=20"`
	AssignmentOfBenefits           string    `hl7:"POS=21"`
	CoordinationOfBenefits         string    `hl7:"POS=22"`
	CoordinationOfBenefitsPriority string    `hl7:"POS=23"`
	NoticeOfAdmissionCode          string    `hl7:"POS=24"`
	NoticeOfAdmissionDate          time.Time `hl7:"POS=25;ATR=date"`
	ReportOfEligibilityCode        string    `hl7:"POS=26"`
	ReportOfEligibilityDate        time.Time `hl7:"POS=27;ATR=date"`
	ReleaseInformationCode         string    `hl7:"POS=28"`
	PreAdmitCertificationPac       string    `hl7:"POS=29"`
	VerificationDateTime           time.Time `hl7:"POS=30"`
	VerificationBy                 CN_PERSON `hl7:"POS=31"`
	TypeOfAgreementCode            string    `hl7:"POS=32"`
	BillingStatus                  string    `hl7:"POS=33"`
	LifetimeReserveDays            *int      `hl7:"POS=34"`
	DelayBeforeLifetimeReserveDays *int      `hl7:"POS=35"`
	CompanyPlanCode                string    `hl7:"POS=36"`
	PolicyNumber                   string    `hl7:"POS=37"`
	PolicyDeductible               *int      `hl7:"POS=38"`
	PolicyLimitAmount              *int      `hl7:"POS=39"`
	PolicyLimitDays                *int      `hl7:"POS=40"`
	RoomRateSemiPrivate            *int      `hl7:"POS=41"`
	RoomRatePrivate                *int      `hl7:"POS=42"`
	InsuredsEmploymentStatus       CE        `hl7:"POS=43"`
	InsuredsSex                    string    `hl7:"POS=44"`
	InsuredsEmployerAddress        AD        `hl7:"POS=45"`
	VerificationStatus             string    `hl7:"POS=46"`
	PriorInsurancePlanID           string    `hl7:"POS=47"`
}
