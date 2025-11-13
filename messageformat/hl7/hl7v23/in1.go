package hl7v23

import "time"

// HL7 v2.3 - IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/IN1
type IN1 struct {
	SetID                          int       `hl7:"POS=2;ATR=sequence" json:"setID,omitempty"`
	PlanID                         CE        `hl7:"POS=3" json:"planID,omitempty"`
	CompanyID                      []CX      `hl7:"POS=4" json:"companyID,omitempty"`
	CompanyName                    []XON     `hl7:"POS=5" json:"companyName,omitempty"`
	CompanyAddress                 []XAD     `hl7:"POS=6" json:"companyAddress,omitempty"`
	ContactPerson                  []XPN     `hl7:"POS=7" json:"contactPerson,omitempty"`
	PhoneNumber                    []XTN     `hl7:"POS=8" json:"phoneNumber,omitempty"`
	GroupNumber                    string    `hl7:"POS=9" json:"groupNumber,omitempty"`
	GroupName                      []XON     `hl7:"POS=10" json:"groupName,omitempty"`
	GroupEmployerID                []CX      `hl7:"POS=11" json:"groupEmployerID,omitempty"`
	GroupEmployerName              []XON     `hl7:"POS=12" json:"groupEmployerName,omitempty"`
	PlanEffectiveDate              time.Time `hl7:"POS=13;ATR=date" json:"planEffectiveDate,omitempty"`
	PlanExpirationDate             time.Time `hl7:"POS=14;ATR=date" json:"planExpirationDate,omitempty"`
	AuthorizationInformation       CM_AUI    `hl7:"POS=15" json:"authorizationInformation,omitempty"`
	PlanType                       string    `hl7:"POS=16" json:"planType,omitempty"`
	NameOfInsured                  []XPN     `hl7:"POS=17" json:"nameOfInsured,omitempty"`
	InsuredRelationshipToPatient   string    `hl7:"POS=18" json:"insuredRelationshipToPatient,omitempty"`
	InsuredDateOfBirth             time.Time `hl7:"POS=19" json:"insuredDateOfBirth,omitempty"`
	InsuredAddress                 []XAD     `hl7:"POS=20" json:"insuredAddress,omitempty"`
	AssignmentOfBenefits           string    `hl7:"POS=21" json:"assignmentOfBenefits,omitempty"`
	CoordinationOfBenefits         string    `hl7:"POS=22" json:"coordinationOfBenefits,omitempty"`
	CoordinationOfBenefitsPriority string    `hl7:"POS=23" json:"coordinationOfBenefitsPriority,omitempty"`
	NoticeOfAdmissionCode          string    `hl7:"POS=24" json:"noticeOfAdmissionCode,omitempty"`
	NoticeOfAdmissionDate          time.Time `hl7:"POS=25;ATR=date" json:"noticeOfAdmissionDate,omitempty"`
	ReportOfEligibilityCode        string    `hl7:"POS=26" json:"reportOfEligibilityCode,omitempty"`
	ReportOfEligibilityDate        time.Time `hl7:"POS=27;ATR=date" json:"reportOfEligibilityDate,omitempty"`
	ReleaseInformationCode         string    `hl7:"POS=28" json:"releaseInformationCode,omitempty"`
	PreAdmitCert                   string    `hl7:"POS=29" json:"preAdmitCert,omitempty"`
	VerificationDateTime           time.Time `hl7:"POS=30" json:"verificationDateTime,omitempty"`
	VerificationBy                 XCN       `hl7:"POS=31" json:"verificationBy,omitempty"`
	TypeOfAgreementCode            string    `hl7:"POS=32" json:"typeOfAgreementCode,omitempty"`
	BillingStatus                  string    `hl7:"POS=33" json:"billingStatus,omitempty"`
	LifetimeReserveDays            *float32  `hl7:"POS=34" json:"lifetimeReserveDays,omitempty"`
	DelayBeforeLifetimeReserveDays *float32  `hl7:"POS=35" json:"delayBeforeLifetimeReserveDays,omitempty"`
	CompanyPlanCode                string    `hl7:"POS=36" json:"companyPlanCode,omitempty"`
	PolicyNumber                   string    `hl7:"POS=37" json:"policyNumber,omitempty"`
	PolicyDeductible               CP        `hl7:"POS=38" json:"policyDeductible,omitempty"`
	PolicyLimitAmount              CP        `hl7:"POS=39" json:"policyLimitAmount,omitempty"`
	PolicyLimitDays                *float32  `hl7:"POS=40" json:"policyLimitDays,omitempty"`
	RoomRateSemiPrivate            CP        `hl7:"POS=41" json:"roomRateSemiPrivate,omitempty"`
	RoomRatePrivate                CP        `hl7:"POS=42" json:"roomRatePrivate,omitempty"`
	InsuredEmploymentStatus        CE        `hl7:"POS=43" json:"insuredEmploymentStatus,omitempty"`
	InsuredSex                     string    `hl7:"POS=44" json:"insuredSex,omitempty"`
	InsuredEmployerAddress         []XAD     `hl7:"POS=45" json:"insuredEmployerAddress,omitempty"`
	VerificationStatus             string    `hl7:"POS=46" json:"verificationStatus,omitempty"`
	PriorInsurancePlanID           string    `hl7:"POS=47" json:"priorInsurancePlanID,omitempty"`
	CoverageType                   string    `hl7:"POS=48" json:"coverageType,omitempty"`
	Handicap                       string    `hl7:"POS=49" json:"handicap,omitempty"`
	InsuredIDNumber                []CX      `hl7:"POS=50" json:"insuredIDNumber,omitempty"`
}
