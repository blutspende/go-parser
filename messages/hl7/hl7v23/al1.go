package hl7v23

import "time"

type AL1 struct {
	SetId              string    `hl7:"1" json:"setId,omitempty"`
	AllergyType        string    `hl7:"2" json:"allergyType,omitempty"`
	AllergyCode        CE        `hl7:"3" json:"allergyCode,omitempty"`
	AllergySeverity    string    `hl7:"4" json:"allergySeverity,omitempty"`
	AllergyReaction    string    `hl7:"5" json:"allergyReaction,omitempty"`
	IdentificationDate time.Time `hl7:"6,shortdate" json:"identificationDate,omitempty"`
}
