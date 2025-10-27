package hl7v23

import "time"

type IN3 struct {
	SetID                              string    `hl7:"POS=2" json:"setID,omitempty"`
	CertificationNumber                CX        `hl7:"POS=3" json:"certificationNumber,omitempty"`
	CertifiedBy                        []XCN     `hl7:"POS=4" json:"certifiedBy,omitempty"`
	CertificationRequired              string    `hl7:"POS=5" json:"certificationRequired,omitempty"`
	Penalty                            CM_PEN    `hl7:"POS=6" json:"penalty,omitempty"`
	CertificationDateTime              time.Time `hl7:"POS=7" json:"certificationDateTime,omitempty"`
	CertificationModifyDateTime        time.Time `hl7:"POS=8" json:"certificationModifyDateTime,omitempty"`
	Operator                           []XCN     `hl7:"POS=9" json:"operator,omitempty"`
	CertificationBeginDate             time.Time `hl7:"POS=10;ATR=date" json:"certificationBeginDate,omitempty"`
	CertificationEndDate               time.Time `hl7:"POS=11;ATR=date" json:"certificationEndDate,omitempty"`
	Days                               CM_DTN    `hl7:"POS=12" json:"days,omitempty"`
	NonConcurCode                      CE        `hl7:"POS=13" json:"nonConcurCode,omitempty"`
	NonConcurEffectiveDateTime         time.Time `hl7:"POS=14" json:"nonConcurEffectiveDateTime,omitempty"`
	PhysicianReviewer                  []XCN     `hl7:"POS=15" json:"physicianReviewer,omitempty"`
	CertificationContact               string    `hl7:"POS=16" json:"certificationContact,omitempty"`
	CertificationContactPhoneNumber    []XTN     `hl7:"POS=17" json:"certificationContactPhoneNumber,omitempty"`
	AppealReason                       CE        `hl7:"POS=18" json:"appealReason,omitempty"`
	CertificationAgency                CE        `hl7:"POS=19" json:"certificationAgency,omitempty"`
	CertificationAgencyPhoneNumber     []XTN     `hl7:"POS=20" json:"certificationAgencyPhoneNumber,omitempty"`
	PreCertificationRequired           []CM_PCF  `hl7:"POS=21" json:"preCertificationRequired,omitempty"`
	CaseManager                        string    `hl7:"POS=22" json:"caseManager,omitempty"`
	SecondOpinionDate                  time.Time `hl7:"POS=23" json:"secondOpinionDate,omitempty"`
	SecondOpinionStatus                string    `hl7:"POS=24" json:"secondOpinionStatus,omitempty"`
	SecondOpinionDocumentationReceived []string  `hl7:"POS=25" json:"secondOpinionDocumentationReceived,omitempty"`
	SecondOptionPhysician              []XCN     `hl7:"POS=26" json:"secondOptionPhysician,omitempty"`
}
