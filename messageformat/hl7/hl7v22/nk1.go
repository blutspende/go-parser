package hl7v22

import "time"

// NK1 - Next Of Kin
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/NK1
type NK1 struct {
	SetID                   int         `hl7:"POS=2;ATR=sequence"`
	Name                    PN          `hl7:"POS=3"`
	Relationship            CE          `hl7:"POS=4"`
	Address                 AD          `hl7:"POS=5"`
	PhoneNumber             []string    `hl7:"POS=6"`
	BusinessPhoneNumber     string      `hl7:"POS=7"`
	ContactRole             CE          `hl7:"POS=8"`
	StartDate               time.Time   `hl7:"POS=9;ATR=date"`
	EndDate                 time.Time   `hl7:"POS=10;ATR=date"`
	NextOfKin               string      `hl7:"POS=11"`
	NextOfKinJobCodeClass   CM_JOB_CODE `hl7:"POS=12"`
	NextOfKinEmployeeNumber string      `hl7:"POS=13"`
	OrganizationName        string      `hl7:"POS=14"`
}
