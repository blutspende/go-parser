package hl7v25

import "time"

// IAM - Patient Adverse Reaction Information
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/IAM
type IAM struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	AllergenTypeCode CE `hl7:"POS=3"`
	AllergenCodeMnemonicDescription CE `hl7:"POS=4"`
	AllergySeverityCode CE `hl7:"POS=5"`
	AllergyReactionCode []string `hl7:"POS=6"`
	AllergyActionCode CNE `hl7:"POS=7"`
	AllergyUniqueIdentifier EI `hl7:"POS=8"`
	ActionReason string `hl7:"POS=9"`
	SensitivityToCausativeAgentCode CE `hl7:"POS=10"`
	AllergenGroupCodeMnemonicDescription CE `hl7:"POS=11"`
	OnsetDate time.Time `hl7:"POS=12;ATR=date"`
	OnsetDateText string `hl7:"POS=13"`
	ReportedDateTime time.Time `hl7:"POS=14"`
	ReportedBy XPN `hl7:"POS=15"`
	RelationshipToPatientCode CE `hl7:"POS=16"`
	AlertDeviceCode CE `hl7:"POS=17"`
	AllergyClinicalStatusCode CE `hl7:"POS=18"`
	StatusedByPerson XCN `hl7:"POS=19"`
	StatusedByOrganization XON `hl7:"POS=20"`
	StatusedAtDateTime time.Time `hl7:"POS=21"`
}

