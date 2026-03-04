package hl7v22

import "time"

// GT1 - Guarantor
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/GT1
type GT1 struct {
	SetID                         int           `hl7:"POS=2;ATR=sequence"`
	GuarantorNumber               COMP_ID_DIGIT `hl7:"POS=3"`
	GuarantorName                 PN            `hl7:"POS=4"`
	GuarantorSpouseName           PN            `hl7:"POS=5"`
	GuarantorAddress              AD            `hl7:"POS=6"`
	GuarantorPhoneNumberHome      []string      `hl7:"POS=7"`
	GuarantorPhoneNumberBusiness  []string      `hl7:"POS=8"`
	GuarantorDateOfBirth          time.Time     `hl7:"POS=9;ATR=date"`
	GuarantorSex                  string        `hl7:"POS=10"`
	GuarantorType                 string        `hl7:"POS=11"`
	GuarantorRelationship         string        `hl7:"POS=12"`
	GuarantorSocialSecurityNumber string        `hl7:"POS=13"`
	GuarantorDateBegin            time.Time     `hl7:"POS=14;ATR=date"`
	GuarantorDateEnd              time.Time     `hl7:"POS=15;ATR=date"`
	GuarantorPriority             *int          `hl7:"POS=16"`
	GuarantorEmployerName         string        `hl7:"POS=17"`
	GuarantorEmployerAddress      AD            `hl7:"POS=18"`
	GuarantorEmployPhoneNumber    []string      `hl7:"POS=19"`
	GuarantorEmployeeIDNumber     string        `hl7:"POS=20"`
	GuarantorEmploymentStatus     string        `hl7:"POS=21"`
	GuarantorOrganization         string        `hl7:"POS=22"`
}
