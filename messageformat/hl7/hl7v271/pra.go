package hl7v271

import "time"

// PRA - Practitioner Detail
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/PRA
type PRA struct {
	PrimaryKeyValuePra CWE `hl7:"POS=2"`
	PractitionerGroup []CWE `hl7:"POS=3"`
	PractitionerCategory []CWE `hl7:"POS=4"`
	ProviderBilling string `hl7:"POS=5"`
	Specialty []SPD `hl7:"POS=6"`
	PractitionerIDNumbers []PLN `hl7:"POS=7"`
	Privileges []PIP `hl7:"POS=8"`
	DateEnteredPractice time.Time `hl7:"POS=9;ATR=date"`
	Institution CWE `hl7:"POS=10"`
	DateLeftPractice time.Time `hl7:"POS=11;ATR=date"`
	GovernmentReimbursementBillingEligibility []CWE `hl7:"POS=12"`
	SetID int `hl7:"POS=13;ATR=sequence"`
}

