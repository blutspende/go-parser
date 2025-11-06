package hl7v23

import "time"

type AL1 struct {
	SetId              string    `hl7:"POS=2" json:"setId,omitempty"`
	AllergyType        string    `hl7:"POS=3" json:"allergyType,omitempty"`
	AllergyCode        CE        `hl7:"POS=4" json:"allergyCode,omitempty"`
	AllergySeverity    string    `hl7:"POS=5" json:"allergySeverity,omitempty"`
	AllergyReaction    string    `hl7:"POS=6" json:"allergyReaction,omitempty"`
	IdentificationDate time.Time `hl7:"POS=7;ATR=date" json:"identificationDate,omitempty"`
}
