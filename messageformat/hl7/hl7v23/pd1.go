package hl7v23

// PD1 - Patient Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PD1
type PD1 struct {
	LivingDependency []string `hl7:"POS=2"`
	LivingArrangement string `hl7:"POS=3"`
	PatientPrimaryFacility []XON `hl7:"POS=4"`
	PatientPrimaryCareProviderNameIDNo []XCN `hl7:"POS=5"`
	StudentIndicator string `hl7:"POS=6"`
	Handicap string `hl7:"POS=7"`
	LivingWill string `hl7:"POS=8"`
	OrganDonor string `hl7:"POS=9"`
	SeparateBill string `hl7:"POS=10"`
	DuplicatePatient []CX `hl7:"POS=11"`
	PublicityIndicator CE `hl7:"POS=12"`
	ProtectionIndicator string `hl7:"POS=13"`
}

