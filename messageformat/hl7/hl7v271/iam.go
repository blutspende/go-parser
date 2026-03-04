package hl7v271

import "time"

// IAM - Patient Adverse Reaction Information
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/IAM
type IAM struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	AllergenTypeCode CWE `hl7:"POS=3"`
	AllergenCodeMnemonicDescription CWE `hl7:"POS=4"`
	AllergySeverityCode CWE `hl7:"POS=5"`
	AllergyReactionCode []string `hl7:"POS=6"`
	AllergyActionCode CNE `hl7:"POS=7"`
	AllergyUniqueIdentifier EI `hl7:"POS=8"`
	ActionReason string `hl7:"POS=9"`
	SensitivityToCausativeAgentCode CWE `hl7:"POS=10"`
	AllergenGroupCodeMnemonicDescription CWE `hl7:"POS=11"`
	OnsetDate time.Time `hl7:"POS=12;ATR=date"`
	OnsetDateText string `hl7:"POS=13"`
	ReportedDateTime time.Time `hl7:"POS=14"`
	ReportedBy XPN `hl7:"POS=15"`
	RelationshipToPatientCode CWE `hl7:"POS=16"`
	AlertDeviceCode CWE `hl7:"POS=17"`
	AllergyClinicalStatusCode CWE `hl7:"POS=18"`
	StatusedByPerson XCN `hl7:"POS=19"`
	StatusedByOrganization XON `hl7:"POS=20"`
	StatusedAtDateTime time.Time `hl7:"POS=21"`
	InactivatedByPerson XCN `hl7:"POS=22"`
	InactivatedDateTime time.Time `hl7:"POS=23"`
	InitiallyRecordedByPerson XCN `hl7:"POS=24"`
	InitiallyRecordedDateTime time.Time `hl7:"POS=25"`
	ModifiedByPerson XCN `hl7:"POS=26"`
	ModifiedDateTime time.Time `hl7:"POS=27"`
	ClinicianIdentifiedCode CWE `hl7:"POS=28"`
	InitiallyRecordedByOrganization XON `hl7:"POS=29"`
	ModifiedByOrganization XON `hl7:"POS=30"`
	InactivatedByOrganization XON `hl7:"POS=31"`
}

