package hl7v24

import "time"

// RXG - Pharmacy/Treatment Give
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/RXG
type RXG struct {
	GiveSubIDCounter *int `hl7:"POS=2"`
	DispenseSubIDCounter *int `hl7:"POS=3"`
	QuantityTiming TQ `hl7:"POS=4"`
	GiveCode CE `hl7:"POS=5"`
	GiveAmountMinimum *int `hl7:"POS=6"`
	GiveAmountMaximum *int `hl7:"POS=7"`
	GiveUnits CE `hl7:"POS=8"`
	GiveDosageForm CE `hl7:"POS=9"`
	AdministrationNotes []CE `hl7:"POS=10"`
	SubstitutionStatus string `hl7:"POS=11"`
	DispenseToLocation LA2 `hl7:"POS=12"`
	NeedsHumanReview string `hl7:"POS=13"`
	PharmacyTreatmentSuppliersSpecialAdministrationInstructions []CE `hl7:"POS=14"`
	GivePer string `hl7:"POS=15"`
	GiveRateAmount string `hl7:"POS=16"`
	GiveRateUnits CE `hl7:"POS=17"`
	GiveStrength *int `hl7:"POS=18"`
	GiveStrengthUnits CE `hl7:"POS=19"`
	SubstanceLotNumber []string `hl7:"POS=20"`
	SubstanceExpirationDate []time.Time `hl7:"POS=21"`
	SubstanceManufacturerName []CE `hl7:"POS=22"`
	Indication []CE `hl7:"POS=23"`
}

