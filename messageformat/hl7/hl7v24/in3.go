package hl7v24

import "time"

// HL7 v2.4 - IN3 - Insurance Additional Information, Certification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/IN3
type IN3 struct {
	SetID                              int       `hl7:"POS=2;ATR=sequence" json:"SetID,omitempty"`
	CertificationNumber                CX        `hl7:"POS=3" json:"CertificationNumber,omitempty"`
	CertifiedBy                        []XCN     `hl7:"POS=4" json:"CertifiedBy,omitempty"`
	CertificationRequired              string    `hl7:"POS=5" json:"CertificationRequired,omitempty"`
	Penalty                            MOP       `hl7:"POS=6" json:"Penalty,omitempty"`
	CertificationDateTime              time.Time `hl7:"POS=7" json:"CertificationDateTime,omitempty"`
	CertificationModifyDateTime        time.Time `hl7:"POS=8" json:"CertificationModifyDateTime,omitempty"`
	Operator                           []XCN     `hl7:"POS=9" json:"Operator,omitempty"`
	CertificationBeginDate             time.Time `hl7:"POS=10;ATR=date" json:"CertificationBeginDate,omitempty"`
	CertificationEndDate               time.Time `hl7:"POS=11;ATR=date" json:"CertificationEndDate,omitempty"`
	Days                               DTN       `hl7:"POS=12" json:"Days,omitempty"`
	NonConcurCode                      CE        `hl7:"POS=13" json:"NonConcurCode,omitempty"`
	NonConcurEffectiveDateTime         time.Time `hl7:"POS=14" json:"NonConcurEffectiveDateTime,omitempty"`
	PhysicianReviewer                  []XCN     `hl7:"POS=15" json:"PhysicianReviewer,omitempty"`
	CertificationContact               string    `hl7:"POS=16" json:"CertificationContact,omitempty"`
	CertificationContactPhoneNumber    []XTN     `hl7:"POS=17" json:"CertificationContactPhoneNumber,omitempty"`
	AppealReason                       CE        `hl7:"POS=18" json:"AppealReason,omitempty"`
	CertificationAgency                CE        `hl7:"POS=19" json:"CertificationAgency,omitempty"`
	CertificationAgencyPhoneNumber     []XTN     `hl7:"POS=20" json:"CertificationAgencyPhoneNumber,omitempty"`
	PreCertificationRequired           []PCF     `hl7:"POS=21" json:"PreCertificationRequired,omitempty"`
	CaseManager                        string    `hl7:"POS=22" json:"CaseManager,omitempty"`
	SecondOpinionDate                  time.Time `hl7:"POS=23;ATR=date" json:"SecondOpinionDate,omitempty"`
	SecondOpinionStatus                string    `hl7:"POS=24" json:"SecondOpinionStatus,omitempty"`
	SecondOpinionDocumentationReceived []string  `hl7:"POS=25" json:"SecondOpinionDocumentationReceived,omitempty"`
	SecondOptionPhysician              []XCN     `hl7:"POS=26" json:"SecondOptionPhysician,omitempty"`
}
