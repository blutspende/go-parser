package hl7v28

import "time"

// PRD - Provider Data
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/PRD
type PRD struct {
	ProviderRole []CWE `hl7:"POS=2"`
	ProviderName []XPN `hl7:"POS=3"`
	ProviderAddress []XAD `hl7:"POS=4"`
	ProviderLocation PL `hl7:"POS=5"`
	ProviderCommunicationInformation []XTN `hl7:"POS=6"`
	PreferredMethodOfContact CWE `hl7:"POS=7"`
	ProviderIdentifiers []PLN `hl7:"POS=8"`
	EffectiveStartDateOfProviderRole time.Time `hl7:"POS=9"`
	EffectiveEndDateOfProviderRole []time.Time `hl7:"POS=10"`
	ProviderOrganizationNameAndIdentifier XON `hl7:"POS=11"`
	ProviderOrganizationAddress []XAD `hl7:"POS=12"`
	ProviderOrganizationLocationInformation []PL `hl7:"POS=13"`
	ProviderOrganizationCommunicationInformation []XTN `hl7:"POS=14"`
	ProviderOrganizationMethodOfContact CWE `hl7:"POS=15"`
}

