package hl7v26

import "time"

// PD1 - Patient Additional Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/PD1
type PD1 struct {
	LivingDependency []string `hl7:"POS=2"`
	LivingArrangement string `hl7:"POS=3"`
	PatientPrimaryFacility []XON `hl7:"POS=4"`
	PatientPrimaryCareProviderNameAndIDNo []XCN `hl7:"POS=5"`
	StudentIndicator string `hl7:"POS=6"`
	Handicap string `hl7:"POS=7"`
	LivingWillCode string `hl7:"POS=8"`
	OrganDonorCode string `hl7:"POS=9"`
	SeparateBill string `hl7:"POS=10"`
	DuplicatePatient []CX `hl7:"POS=11"`
	PublicityCode CWE `hl7:"POS=12"`
	ProtectionIndicator string `hl7:"POS=13"`
	ProtectionIndicatorEffectiveDate time.Time `hl7:"POS=14;ATR=date"`
	PlaceOfWorship []XON `hl7:"POS=15"`
	AdvanceDirectiveCode []CWE `hl7:"POS=16"`
	ImmunizationRegistryStatus string `hl7:"POS=17"`
	ImmunizationRegistryStatusEffectiveDate time.Time `hl7:"POS=18;ATR=date"`
	PublicityCodeEffectiveDate time.Time `hl7:"POS=19;ATR=date"`
	MilitaryBranch string `hl7:"POS=20"`
	MilitaryRankGrade string `hl7:"POS=21"`
	MilitaryStatus string `hl7:"POS=22"`
	AdvanceDirectiveLastVerifiedDate time.Time `hl7:"POS=23;ATR=date"`
}

