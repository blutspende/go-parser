package hl7v23

import "time"

// HL7 v2.3 - IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/IN1
type IN1 struct {
	SetID                          string    `hl7:"1" json:"setID,omitempty"`
	PlanID                         CE        `hl7:"2" json:"planID,omitempty"`
	CompanyID                      []CX      `hl7:"3" json:"companyID,omitempty"`
	CompanyName                    []XON     `hl7:"4" json:"companyName,omitempty"`
	CompanyAddress                 []XAD     `hl7:"5" json:"companyAddress,omitempty"`
	ContactPerson                  []XPN     `hl7:"6" json:"contactPerson,omitempty"`
	PhoneNumber                    []XTN     `hl7:"7" json:"phoneNumber,omitempty"`
	GroupNumber                    string    `hl7:"8" json:"groupNumber,omitempty"`
	GroupName                      []XON     `hl7:"9" json:"groupName,omitempty"`
	GroupEmployerID                []CX      `hl7:"10" json:"groupEmployerID,omitempty"`
	GroupEmployerName              []XON     `hl7:"11" json:"groupEmployerName,omitempty"`
	PlanEffectiveDate              time.Time `hl7:"12,shortdate" json:"planEffectiveDate,omitempty"`
	PlanExpirationDate             time.Time `hl7:"13,shortdate" json:"planExpirationDate,omitempty"`
	AuthorizationInformation       CM_AUI    `hl7:"14" json:"authorizationInformation,omitempty"`
	PlanType                       string    `hl7:"15" json:"planType,omitempty"`
	NameOfInsured                  []XPN     `hl7:"16" json:"nameOfInsured,omitempty"`
	InsuredRelationshipToPatient   string    `hl7:"17" json:"insuredRelationshipToPatient,omitempty"`
	InsuredDateOfBirth             time.Time `hl7:"18,longdate" json:"insuredDateOfBirth,omitempty"`
	InsuredAddress                 []XAD     `hl7:"19" json:"insuredAddress,omitempty"`
	AssignmentOfBenefits           string    `hl7:"20" json:"assignmentOfBenefits,omitempty"`
	CoordinationOfBenefits         string    `hl7:"21" json:"coordinationOfBenefits,omitempty"`
	CoordinationOfBenefitsPrio     string    `hl7:"22" json:"coordinationOfBenefitsPrio,omitempty"`
	NoticeOfAdmissionCode          string    `hl7:"23" json:"noticeOfAdmissionCode,omitempty"`
	NoticeOfAdmissionDate          time.Time `hl7:"24,shortdate" json:"noticeOfAdmissionDate,omitempty"`
	ReportOfEligibilityCode        string    `hl7:"25" json:"reportOfEligibilityCode,omitempty"`
	ReportOfEligibilityDate        time.Time `hl7:"26,shortdate" json:"reportOfEligibilityDate,omitempty"`
	ReleaseInformationCode         string    `hl7:"27" json:"releaseInformationCode,omitempty"`
	PreAdmitCert                   string    `hl7:"28" json:"preAdmitCert,omitempty"`
	VerificationDateTime           time.Time `hl7:"29,longdate" json:"verificationDateTime,omitempty"`
	VerificationBy                 XCN       `hl7:"30" json:"verificationBy,omitempty"`
	TypeOfAgreementCode            string    `hl7:"31" json:"typeOfAgreementCode,omitempty"`
	BillingStatus                  string    `hl7:"32" json:"billingStatus,omitempty"`
	LifetimeReserveDays            float32   `hl7:"33" json:"lifetimeReserveDays,omitempty"`
	DelayBeforeLifetimeReserveDays float32   `hl7:"34" json:"delayBeforeLifetimeReserveDays,omitempty"`
	CompanyPlanCode                string    `hl7:"35" json:"companyPlanCode,omitempty"`
	PolicyNumber                   string    `hl7:"36" json:"policyNumber,omitempty"`
	PolicyDeductible               CP        `hl7:"37" json:"policyDeductible,omitempty"`
	PolicyLimitAmount              CP        `hl7:"38" json:"policyLimitAmount,omitempty"`
	PolicyLimitDays                float32   `hl7:"39" json:"policyLimitDays,omitempty"`
	RoomRateSemiPrivate            CP        `hl7:"40" json:"roomRateSemiPrivate,omitempty"`
	RoomRatePrivate                CP        `hl7:"41" json:"roomRatePrivate,omitempty"`
	InsuredEmploymentStatus        CE        `hl7:"42" json:"insuredEmploymentStatus,omitempty"`
	InsuredSex                     string    `hl7:"43" json:"insuredSex,omitempty"`
	InsuredEmployerAddress         []XAD     `hl7:"44" json:"insuredEmployerAddress,omitempty"`
	VerificationStatus             string    `hl7:"45" json:"verificationStatus,omitempty"`
	PriorInsurancePlanID           string    `hl7:"46" json:"priorInsurancePlanID,omitempty"`
	CoverageType                   string    `hl7:"47" json:"coverageType,omitempty"`
	Handicap                       string    `hl7:"48" json:"handicap,omitempty"`
	InsuredIDNumber                []CX      `hl7:"49" json:"insuredIDNumber,omitempty"`
}
