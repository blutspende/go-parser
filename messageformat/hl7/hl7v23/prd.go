package hl7v23

import "time"

// PRD - Provider Data
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PRD
type PRD struct {
	Role []CE `hl7:"POS=2"`
	ProviderName []XPN `hl7:"POS=3"`
	ProviderAddress XAD `hl7:"POS=4"`
	ProviderLocation PL `hl7:"POS=5"`
	ProviderCommunicationInformation []XTN `hl7:"POS=6"`
	PreferredMethodOfContact CE `hl7:"POS=7"`
	ProviderIdentifiers []CM_PI `hl7:"POS=8"`
	EffectiveStartDateOfRole time.Time `hl7:"POS=9"`
	EffectiveEndDateOfRole time.Time `hl7:"POS=10"`
}

