package hl7v25

import "time"

// IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/IN1
type IN1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	InsurancePlanID CE `hl7:"POS=3"`
	InsuranceCompanyID []CX `hl7:"POS=4"`
	InsuranceCompanyName []XON `hl7:"POS=5"`
	InsuranceCompanyAddress []XAD `hl7:"POS=6"`
	InsuranceCoContactPerson []XPN `hl7:"POS=7"`
	InsuranceCoPhoneNumber []XTN `hl7:"POS=8"`
	GroupNumber string `hl7:"POS=9"`
	GroupName []XON `hl7:"POS=10"`
	InsuredsGroupEmpID []CX `hl7:"POS=11"`
	InsuredsGroupEmpName []XON `hl7:"POS=12"`
	PlanEffectiveDate time.Time `hl7:"POS=13;ATR=date"`
	PlanExpirationDate time.Time `hl7:"POS=14;ATR=date"`
	AuthorizationInformation AUI `hl7:"POS=15"`
	PlanType string `hl7:"POS=16"`
	NameOfInsured []XPN `hl7:"POS=17"`
	InsuredsRelationshipToPatient CE `hl7:"POS=18"`
	InsuredsDateOfBirth time.Time `hl7:"POS=19"`
	InsuredsAddress []XAD `hl7:"POS=20"`
	AssignmentOfBenefits string `hl7:"POS=21"`
	CoordinationOfBenefits string `hl7:"POS=22"`
	CoordOfBenPriority string `hl7:"POS=23"`
	NoticeOfAdmissionFlag string `hl7:"POS=24"`
	NoticeOfAdmissionDate time.Time `hl7:"POS=25;ATR=date"`
	ReportOfEligibilityFlag string `hl7:"POS=26"`
	ReportOfEligibilityDate time.Time `hl7:"POS=27;ATR=date"`
	ReleaseInformationCode string `hl7:"POS=28"`
	PreAdmitCertPac string `hl7:"POS=29"`
	VerificationDateTime time.Time `hl7:"POS=30"`
	VerificationBy []XCN `hl7:"POS=31"`
	TypeOfAgreementCode string `hl7:"POS=32"`
	BillingStatus string `hl7:"POS=33"`
	LifetimeReserveDays *int `hl7:"POS=34"`
	DelayBeforeLrDay *int `hl7:"POS=35"`
	CompanyPlanCode string `hl7:"POS=36"`
	PolicyNumber string `hl7:"POS=37"`
	PolicyDeductible CP `hl7:"POS=38"`
	PolicyLimitAmount CP `hl7:"POS=39"`
	PolicyLimitDays *int `hl7:"POS=40"`
	RoomRateSemiPrivate CP `hl7:"POS=41"`
	RoomRatePrivate CP `hl7:"POS=42"`
	InsuredsEmploymentStatus CE `hl7:"POS=43"`
	InsuredsAdministrativeSex string `hl7:"POS=44"`
	InsuredsEmployersAddress []XAD `hl7:"POS=45"`
	VerificationStatus string `hl7:"POS=46"`
	PriorInsurancePlanID string `hl7:"POS=47"`
	CoverageType string `hl7:"POS=48"`
	Handicap string `hl7:"POS=49"`
	InsuredsIDNumber []CX `hl7:"POS=50"`
	SignatureCode string `hl7:"POS=51"`
	SignatureCodeDate time.Time `hl7:"POS=52;ATR=date"`
	InsuredsBirthPlace string `hl7:"POS=53"`
	VipIndicator string `hl7:"POS=54"`
}

