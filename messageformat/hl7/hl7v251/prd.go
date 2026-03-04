package hl7v251

import "time"

// PRD - Provider Data
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/PRD
type PRD struct {
	ProviderRole []CE `hl7:"POS=2"`
	ProviderName []XPN `hl7:"POS=3"`
	ProviderAddress []XAD `hl7:"POS=4"`
	ProviderLocation PL `hl7:"POS=5"`
	ProviderCommunicationInformation []XTN `hl7:"POS=6"`
	PreferredMethodOfContact CE `hl7:"POS=7"`
	ProviderIdentifiers []PLN `hl7:"POS=8"`
	EffectiveStartDateOfProviderRole time.Time `hl7:"POS=9"`
	EffectiveEndDateOfProviderRole time.Time `hl7:"POS=10"`
}

