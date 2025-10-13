package hl7v24

import "time"

// HL7 v2.4 - ROL - Role
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ROL
type ROL struct {
	RoleInstanceID       EI        `hl7:"2" json:"RoleInstanceID,omitempty"`
	ActionCode           string    `hl7:"3" json:"ActionCode,omitempty"`
	RoleROL              CE        `hl7:"4" json:"RoleROL,omitempty"`
	RolePerson           []XCN     `hl7:"5" json:"RolePerson,omitempty"`
	RoleBeginDateTime    time.Time `hl7:"6" json:"RoleBeginDateTime,omitempty"`
	RoleEndDateTime      time.Time `hl7:"7" json:"RoleEndDateTime,omitempty"`
	RoleDuration         CE        `hl7:"8" json:"RoleDuration,omitempty"`
	RoleActionReason     CE        `hl7:"9" json:"RoleActionReason,omitempty"`
	ProviderType         []CE      `hl7:"10"json:"ProviderType,omitempty"`
	OrganizationUnitType CE        `hl7:"11" json:"OrganizationUnitType,omitempty"`
	OfficeHomeAddress    []XAD     `hl7:"12" json:"OfficeHomeAddress,omitempty"`
	Phone                []XTN     `hl7:"13" json:"Phone,omitempty"`
}
