package hl7v25

import "time"

// IN3 - Insurance Additional Information, Certification
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/IN3
type IN3 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	CertificationNumber CX `hl7:"POS=3"`
	CertifiedBy []XCN `hl7:"POS=4"`
	CertificationRequired string `hl7:"POS=5"`
	Penalty MOP `hl7:"POS=6"`
	CertificationDateTime time.Time `hl7:"POS=7"`
	CertificationModifyDateTime time.Time `hl7:"POS=8"`
	Operator []XCN `hl7:"POS=9"`
	CertificationBeginDate time.Time `hl7:"POS=10;ATR=date"`
	CertificationEndDate time.Time `hl7:"POS=11;ATR=date"`
	Days DTN `hl7:"POS=12"`
	NonConcurCodeDescription CE `hl7:"POS=13"`
	NonConcurEffectiveDateTime time.Time `hl7:"POS=14"`
	PhysicianReviewer []XCN `hl7:"POS=15"`
	CertificationContact string `hl7:"POS=16"`
	CertificationContactPhoneNumber []XTN `hl7:"POS=17"`
	AppealReason CE `hl7:"POS=18"`
	CertificationAgency CE `hl7:"POS=19"`
	CertificationAgencyPhoneNumber []XTN `hl7:"POS=20"`
	PreCertificationRequirement []ICD `hl7:"POS=21"`
	CaseManager string `hl7:"POS=22"`
	SecondOpinionDate time.Time `hl7:"POS=23;ATR=date"`
	SecondOpinionStatus string `hl7:"POS=24"`
	SecondOpinionDocumentationReceived []string `hl7:"POS=25"`
	SecondOpinionPhysician []XCN `hl7:"POS=26"`
}

