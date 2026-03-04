package hl7v271

import "time"

// ROL - Role
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/ROL
type ROL struct {
	RoleInstanceID EI `hl7:"POS=2"`
	ActionCode string `hl7:"POS=3"`
	RoleRol CWE `hl7:"POS=4"`
	RolePerson []XCN `hl7:"POS=5"`
	RoleBeginDateTime time.Time `hl7:"POS=6"`
	RoleEndDateTime time.Time `hl7:"POS=7"`
	RoleDuration CWE `hl7:"POS=8"`
	RoleActionReason CWE `hl7:"POS=9"`
	ProviderType []CWE `hl7:"POS=10"`
	OrganizationUnitType CWE `hl7:"POS=11"`
	OfficeHomeAddressBirthplace []XAD `hl7:"POS=12"`
	Phone []XTN `hl7:"POS=13"`
	PersonsLocation PL `hl7:"POS=14"`
	Organization XON `hl7:"POS=15"`
}

