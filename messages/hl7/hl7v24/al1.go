package hl7v24

import (
	"time"
)

// HL7 v2.4 - AL1 - Patient allergy information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/AL1
type AL1 struct {
	SetId               CE        `hl7:"1" json:"SetId,omitempty"`
	AllergenTypeCode    CE        `hl7:"2" json:"AllergenTypeCode,omitempty"`
	AllergenCode        CE        `hl7:"3" json:"AllergenCode,omitempty"`
	AllergySeverityCode CE        `hl7:"4" json:"AllergySeverityCode,omitempty"`
	AllergyReactionCode string    `hl7:"5" json:"AllergyReactionCode,omitempty"`
	IdentificationDate  time.Time `hl7:"6,shortdate" json:"IdentificationDate,omitempty"`
}
