package hl7v24

import "time"

// HL7 v2.4 - IN3 - Insurance Additional Information, Certification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/IN3
type IN3 struct {
	SetID                              string    `hl7:"2" json:"SetID,omitempty"`
	CertificationNumber                CX        `hl7:"3" json:"CertificationNumber,omitempty"`
	CertifiedBy                        []XCN     `hl7:"4" json:"CertifiedBy,omitempty"`
	CertificationRequired              string    `hl7:"5" json:"CertificationRequired,omitempty"`
	Penalty                            MOP       `hl7:"6" json:"Penalty,omitempty"`
	CertificationDateTime              time.Time `hl7:"7,longdate" json:"CertificationDateTime,omitempty"`
	CertificationModifyDateTime        time.Time `hl7:"8,longdate" json:"CertificationModifyDateTime,omitempty"`
	Operator                           []XCN     `hl7:"9" json:"Operator,omitempty"`
	CertificationBeginDate             time.Time `hl7:"10" json:"CertificationBeginDate,omitempty"`
	CertificationEndDate               time.Time `hl7:"11" json:"CertificationEndDate,omitempty"`
	Days                               DTN       `hl7:"12" json:"Days,omitempty"`
	NonConcurCode                      CE        `hl7:"13" json:"NonConcurCode,omitempty"`
	NonConcurEffectiveDateTime         time.Time `hl7:"14,longdate" json:"NonConcurEffectiveDateTime,omitempty"`
	PhysicianReviewer                  []XCN     `hl7:"15" json:"PhysicianReviewer,omitempty"`
	CertificationContact               string    `hl7:"16" json:"CertificationContact,omitempty"`
	CertificationContactPhoneNumber    []XTN     `hl7:"17" json:"CertificationContactPhoneNumber,omitempty"`
	AppealReason                       CE        `hl7:"18" json:"AppealReason,omitempty"`
	CertificationAgency                CE        `hl7:"19" json:"CertificationAgency,omitempty"`
	CertificationAgencyPhoneNumber     []XTN     `hl7:"20" json:"CertificationAgencyPhoneNumber,omitempty"`
	PreCertificationRequired           []PCF     `hl7:"21" json:"PreCertificationRequired,omitempty"`
	CaseManager                        string    `hl7:"22" json:"CaseManager,omitempty"`
	SecondOpinionDate                  time.Time `hl7:"23" json:"SecondOpinionDate,omitempty"`
	SecondOpinionStatus                string    `hl7:"24" json:"SecondOpinionStatus,omitempty"`
	SecondOpinionDocumentationReceived []string  `hl7:"25" json:"SecondOpinionDocumentationReceived,omitempty"`
	SecondOptionPhysician              []XCN     `hl7:"26" json:"SecondOptionPhysician,omitempty"`
}
