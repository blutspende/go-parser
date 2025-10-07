package hl7v23

import "time"

type IN3 struct {
	SetID                              string    `hl7:"1" json:"setID,omitempty"`
	CertificationNumber                CX        `hl7:"2" json:"certificationNumber,omitempty"`
	CertifiedBy                        []XCN     `hl7:"3" json:"certifiedBy,omitempty"`
	CertificationRequired              string    `hl7:"4" json:"certificationRequired,omitempty"`
	Penalty                            CM_PEN    `hl7:"5" json:"penalty,omitempty"`
	CertificationDateTime              time.Time `hl7:"6,longdate" json:"certificationDateTime,omitempty"`
	CertificationModifyDateTime        time.Time `hl7:"7,longdate" json:"certificationModifyDateTime,omitempty"`
	Operator                           []XCN     `hl7:"8" json:"operator,omitempty"`
	CertificationBeginDate             time.Time `hl7:"9,shortdate" json:"certificationBeginDate,omitempty"`
	CertificationEndDate               time.Time `hl7:"10,shortdate" json:"certificationEndDate,omitempty"`
	Days                               CM_DTN    `hl7:"11" json:"days,omitempty"`
	NonConcurCode                      CE        `hl7:"12" json:"nonConcurCode,omitempty"`
	NonConcurEffectiveDateTime         time.Time `hl7:"13,longdate" json:"nonConcurEffectiveDateTime,omitempty"`
	PhysicianReviewer                  []XCN     `hl7:"14" json:"physicianReviewer,omitempty"`
	CertificationContact               string    `hl7:"15" json:"certificationContact,omitempty"`
	CertificationContactPhoneNumber    []XTN     `hl7:"16" json:"certificationContactPhoneNumber,omitempty"`
	AppealReason                       CE        `hl7:"17" json:"appealReason,omitempty"`
	CertificationAgency                CE        `hl7:"18" json:"certificationAgency,omitempty"`
	CertificationAgencyPhoneNumber     []XTN     `hl7:"19" json:"certificationAgencyPhoneNumber,omitempty"`
	PreCertificationRequired           []CM_PCF  `hl7:"20" json:"preCertificationRequired,omitempty"`
	CaseManager                        string    `hl7:"21" json:"caseManager,omitempty"`
	SecondOpinionDate                  time.Time `hl7:"22,longdate" json:"secondOpinionDate,omitempty"`
	SecondOpinionStatus                string    `hl7:"23" json:"secondOpinionStatus,omitempty"`
	SecondOpinionDocumentationReceived []string  `hl7:"24" json:"secondOpinionDocumentationReceived,omitempty"`
	SecondOptionPhysician              []XCN     `hl7:"25" json:"secondOptionPhysician,omitempty"`
}
