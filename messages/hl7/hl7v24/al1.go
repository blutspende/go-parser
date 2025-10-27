package hl7v24

import (
	"time"
)

// HL7 v2.4 - AL1 - Patient allergy information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/AL1
type AL1 struct {
	SetId               CE        `hl7:"POS=2" json:"SetId,omitempty"`
	AllergenTypeCode    CE        `hl7:"POS=3" json:"AllergenTypeCode,omitempty"`
	AllergenCode        CE        `hl7:"POS=4" json:"AllergenCode,omitempty"`
	AllergySeverityCode CE        `hl7:"POS=5" json:"AllergySeverityCode,omitempty"`
	AllergyReactionCode string    `hl7:"POS=6" json:"AllergyReactionCode,omitempty"`
	IdentificationDate  time.Time `hl7:"POS=7;ATR=date" json:"IdentificationDate,omitempty"`
}
