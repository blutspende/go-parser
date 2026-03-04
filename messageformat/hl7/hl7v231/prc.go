package hl7v231

import "time"

// PRC - Pricing segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/PRC
type PRC struct {
	PrimaryKeyValuePrc CE `hl7:"POS=2"`
	FacilityIDPrc []CE `hl7:"POS=3"`
	Department []CE `hl7:"POS=4"`
	ValidPatientClasses []string `hl7:"POS=5"`
	Price []CP `hl7:"POS=6"`
	Formula []string `hl7:"POS=7"`
	MinimumQuantity *int `hl7:"POS=8"`
	MaximumQuantity *int `hl7:"POS=9"`
	MinimumPrice MO `hl7:"POS=10"`
	MaximumPrice MO `hl7:"POS=11"`
	EffectiveStartDate time.Time `hl7:"POS=12"`
	EffectiveEndDate time.Time `hl7:"POS=13"`
	PriceOverrideFlag string `hl7:"POS=14"`
	BillingCategory []CE `hl7:"POS=15"`
	ChargeableFlag string `hl7:"POS=16"`
	ActiveInactiveFlag string `hl7:"POS=17"`
	Cost MO `hl7:"POS=18"`
	ChargeOnIndicator string `hl7:"POS=19"`
}

