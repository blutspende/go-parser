package hl7v27

// AFF - Professional Affiliation
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/AFF
type AFF struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ProfessionalOrganization XON `hl7:"POS=3"`
	ProfessionalOrganizationAddress XAD `hl7:"POS=4"`
	ProfessionalOrganizationAffiliationDateRange []DR `hl7:"POS=5"`
	ProfessionalAffiliationAdditionalInformation string `hl7:"POS=6"`
}

