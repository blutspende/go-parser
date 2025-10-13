package hl7v23

import "time"

// HL7 v2.3 - IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/IN1
type IN1 struct {
	SetID                          string    `hl7:"2" json:"setID,omitempty"`
	PlanID                         CE        `hl7:"3" json:"planID,omitempty"`
	CompanyID                      []CX      `hl7:"4" json:"companyID,omitempty"`
	CompanyName                    []XON     `hl7:"5" json:"companyName,omitempty"`
	CompanyAddress                 []XAD     `hl7:"6" json:"companyAddress,omitempty"`
	ContactPerson                  []XPN     `hl7:"7" json:"contactPerson,omitempty"`
	PhoneNumber                    []XTN     `hl7:"8" json:"phoneNumber,omitempty"`
	GroupNumber                    string    `hl7:"9" json:"groupNumber,omitempty"`
	GroupName                      []XON     `hl7:"10" json:"groupName,omitempty"`
	GroupEmployerID                []CX      `hl7:"11" json:"groupEmployerID,omitempty"`
	GroupEmployerName              []XON     `hl7:"12" json:"groupEmployerName,omitempty"`
	PlanEffectiveDate              time.Time `hl7:"13" json:"planEffectiveDate,omitempty"`
	PlanExpirationDate             time.Time `hl7:"14" json:"planExpirationDate,omitempty"`
	AuthorizationInformation       CM_AUI    `hl7:"15" json:"authorizationInformation,omitempty"`
	PlanType                       string    `hl7:"16" json:"planType,omitempty"`
	NameOfInsured                  []XPN     `hl7:"17" json:"nameOfInsured,omitempty"`
	InsuredRelationshipToPatient   string    `hl7:"18" json:"insuredRelationshipToPatient,omitempty"`
	InsuredDateOfBirth             time.Time `hl7:"19,longdate" json:"insuredDateOfBirth,omitempty"`
	InsuredAddress                 []XAD     `hl7:"20" json:"insuredAddress,omitempty"`
	AssignmentOfBenefits           string    `hl7:"21" json:"assignmentOfBenefits,omitempty"`
	CoordinationOfBenefits         string    `hl7:"22" json:"coordinationOfBenefits,omitempty"`
	CoordinationOfBenefitsPrio     string    `hl7:"23" json:"coordinationOfBenefitsPrio,omitempty"`
	NoticeOfAdmissionCode          string    `hl7:"24" json:"noticeOfAdmissionCode,omitempty"`
	NoticeOfAdmissionDate          time.Time `hl7:"25" json:"noticeOfAdmissionDate,omitempty"`
	ReportOfEligibilityCode        string    `hl7:"26" json:"reportOfEligibilityCode,omitempty"`
	ReportOfEligibilityDate        time.Time `hl7:"27" json:"reportOfEligibilityDate,omitempty"`
	ReleaseInformationCode         string    `hl7:"28" json:"releaseInformationCode,omitempty"`
	PreAdmitCert                   string    `hl7:"29" json:"preAdmitCert,omitempty"`
	VerificationDateTime           time.Time `hl7:"30,longdate" json:"verificationDateTime,omitempty"`
	VerificationBy                 XCN       `hl7:"31" json:"verificationBy,omitempty"`
	TypeOfAgreementCode            string    `hl7:"32" json:"typeOfAgreementCode,omitempty"`
	BillingStatus                  string    `hl7:"33" json:"billingStatus,omitempty"`
	LifetimeReserveDays            float32   `hl7:"34" json:"lifetimeReserveDays,omitempty"`
	DelayBeforeLifetimeReserveDays float32   `hl7:"35" json:"delayBeforeLifetimeReserveDays,omitempty"`
	CompanyPlanCode                string    `hl7:"36" json:"companyPlanCode,omitempty"`
	PolicyNumber                   string    `hl7:"37" json:"policyNumber,omitempty"`
	PolicyDeductible               CP        `hl7:"38" json:"policyDeductible,omitempty"`
	PolicyLimitAmount              CP        `hl7:"39" json:"policyLimitAmount,omitempty"`
	PolicyLimitDays                float32   `hl7:"40" json:"policyLimitDays,omitempty"`
	RoomRateSemiPrivate            CP        `hl7:"41" json:"roomRateSemiPrivate,omitempty"`
	RoomRatePrivate                CP        `hl7:"42" json:"roomRatePrivate,omitempty"`
	InsuredEmploymentStatus        CE        `hl7:"43" json:"insuredEmploymentStatus,omitempty"`
	InsuredSex                     string    `hl7:"44" json:"insuredSex,omitempty"`
	InsuredEmployerAddress         []XAD     `hl7:"45" json:"insuredEmployerAddress,omitempty"`
	VerificationStatus             string    `hl7:"46" json:"verificationStatus,omitempty"`
	PriorInsurancePlanID           string    `hl7:"47" json:"priorInsurancePlanID,omitempty"`
	CoverageType                   string    `hl7:"48" json:"coverageType,omitempty"`
	Handicap                       string    `hl7:"49" json:"handicap,omitempty"`
	InsuredIDNumber                []CX      `hl7:"50" json:"insuredIDNumber,omitempty"`
}
