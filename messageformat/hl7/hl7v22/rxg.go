package hl7v22

// RXG - Pharmacy Give
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/RXG
type RXG struct {
	GiveSubIDCounter *int `hl7:"POS=2"`
	DispenseSubIDCounter *int `hl7:"POS=3"`
	QuantityTiming TQ `hl7:"POS=4"`
	GiveCode CE `hl7:"POS=5"`
	GiveAmountMinimum *int `hl7:"POS=6"`
	GiveAmountMaximum *int `hl7:"POS=7"`
	GiveUnits CE `hl7:"POS=8"`
	GiveDosageForm CE `hl7:"POS=9"`
	AdministrationNotes []string `hl7:"POS=10"`
	SubstitutionStatus string `hl7:"POS=11"`
	DispenseToLocation CM_LA1 `hl7:"POS=12"`
	NeedsHumanReview string `hl7:"POS=13"`
	PharmacySpecialAdministrationInstructions []CE `hl7:"POS=14"`
	GivePerTimeUnit string `hl7:"POS=15"`
	GiveRateAmount CE `hl7:"POS=16"`
	GiveRateUnits CE `hl7:"POS=17"`
}

