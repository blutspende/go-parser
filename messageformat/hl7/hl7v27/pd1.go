package hl7v27

import "time"

// PD1 - Patient Additional Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/PD1
type PD1 struct {
	LivingDependency []CWE `hl7:"POS=2"`
	LivingArrangement CWE `hl7:"POS=3"`
	PatientPrimaryFacility []XON `hl7:"POS=4"`
	StudentIndicator CWE `hl7:"POS=6"`
	Handicap CWE `hl7:"POS=7"`
	LivingWillCode CWE `hl7:"POS=8"`
	OrganDonorCode CWE `hl7:"POS=9"`
	SeparateBill string `hl7:"POS=10"`
	DuplicatePatient []CX `hl7:"POS=11"`
	PublicityCode CWE `hl7:"POS=12"`
	ProtectionIndicator string `hl7:"POS=13"`
	ProtectionIndicatorEffectiveDate time.Time `hl7:"POS=14;ATR=date"`
	PlaceOfWorship []XON `hl7:"POS=15"`
	AdvanceDirectiveCode []CWE `hl7:"POS=16"`
	ImmunizationRegistryStatus CWE `hl7:"POS=17"`
	ImmunizationRegistryStatusEffectiveDate time.Time `hl7:"POS=18;ATR=date"`
	PublicityCodeEffectiveDate time.Time `hl7:"POS=19;ATR=date"`
	MilitaryBranch CWE `hl7:"POS=20"`
	MilitaryRankGrade CWE `hl7:"POS=21"`
	MilitaryStatus CWE `hl7:"POS=22"`
	AdvanceDirectiveLastVerifiedDate time.Time `hl7:"POS=23;ATR=date"`
}

