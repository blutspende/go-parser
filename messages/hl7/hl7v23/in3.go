package hl7v23

import "time"

type IN3 struct {
	SetID                              string    `hl7:"2" json:"setID,omitempty"`
	CertificationNumber                CX        `hl7:"3" json:"certificationNumber,omitempty"`
	CertifiedBy                        []XCN     `hl7:"4" json:"certifiedBy,omitempty"`
	CertificationRequired              string    `hl7:"5" json:"certificationRequired,omitempty"`
	Penalty                            CM_PEN    `hl7:"6" json:"penalty,omitempty"`
	CertificationDateTime              time.Time `hl7:"7,longdate" json:"certificationDateTime,omitempty"`
	CertificationModifyDateTime        time.Time `hl7:"8,longdate" json:"certificationModifyDateTime,omitempty"`
	Operator                           []XCN     `hl7:"9" json:"operator,omitempty"`
	CertificationBeginDate             time.Time `hl7:"10"json:"certificationBeginDate,omitempty"`
	CertificationEndDate               time.Time `hl7:"11" json:"certificationEndDate,omitempty"`
	Days                               CM_DTN    `hl7:"12" json:"days,omitempty"`
	NonConcurCode                      CE        `hl7:"13" json:"nonConcurCode,omitempty"`
	NonConcurEffectiveDateTime         time.Time `hl7:"14,longdate" json:"nonConcurEffectiveDateTime,omitempty"`
	PhysicianReviewer                  []XCN     `hl7:"15" json:"physicianReviewer,omitempty"`
	CertificationContact               string    `hl7:"16" json:"certificationContact,omitempty"`
	CertificationContactPhoneNumber    []XTN     `hl7:"17" json:"certificationContactPhoneNumber,omitempty"`
	AppealReason                       CE        `hl7:"18" json:"appealReason,omitempty"`
	CertificationAgency                CE        `hl7:"19" json:"certificationAgency,omitempty"`
	CertificationAgencyPhoneNumber     []XTN     `hl7:"20" json:"certificationAgencyPhoneNumber,omitempty"`
	PreCertificationRequired           []CM_PCF  `hl7:"21" json:"preCertificationRequired,omitempty"`
	CaseManager                        string    `hl7:"22" json:"caseManager,omitempty"`
	SecondOpinionDate                  time.Time `hl7:"23,longdate" json:"secondOpinionDate,omitempty"`
	SecondOpinionStatus                string    `hl7:"24" json:"secondOpinionStatus,omitempty"`
	SecondOpinionDocumentationReceived []string  `hl7:"25" json:"secondOpinionDocumentationReceived,omitempty"`
	SecondOptionPhysician              []XCN     `hl7:"26" json:"secondOptionPhysician,omitempty"`
}
