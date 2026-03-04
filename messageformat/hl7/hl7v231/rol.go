package hl7v231

import "time"

// ROL - Role
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/ROL
type ROL struct {
	RoleInstanceID EI `hl7:"POS=2"`
	ActionCode string `hl7:"POS=3"`
	RoleRol CE `hl7:"POS=4"`
	RolePerson []XCN `hl7:"POS=5"`
	RoleBeginDateTime time.Time `hl7:"POS=6"`
	RoleEndDateTime time.Time `hl7:"POS=7"`
	RoleDuration CE `hl7:"POS=8"`
	RoleActionReason CE `hl7:"POS=9"`
}

