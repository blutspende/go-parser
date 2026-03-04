package hl7v22

import "time"

// STF - Staff Identification Segment
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/STF
type STF struct {
	StfPrimaryKeyValue CE `hl7:"POS=2"`
	StaffIdcode []CE `hl7:"POS=3"`
	StaffName PN `hl7:"POS=4"`
	StaffType []string `hl7:"POS=5"`
	Sex string `hl7:"POS=6"`
	DateOfBirth time.Time `hl7:"POS=7"`
	ActiveInactive string `hl7:"POS=8"`
	Department []CE `hl7:"POS=9"`
	Service []CE `hl7:"POS=10"`
	Phone []string `hl7:"POS=11"`
	OfficeHomeAddress []AD `hl7:"POS=12"`
	ActivationDate []CM_DIN `hl7:"POS=13"`
	InactivationDate []CM_DIN `hl7:"POS=14"`
	BackupPersonID []CE `hl7:"POS=15"`
	EMailAddress []string `hl7:"POS=16"`
	PreferredMethodOfContact string `hl7:"POS=17"`
}

