package hl7v23

import "time"

// RXD - Pharmacy dispense segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RXD
type RXD struct {
	DispenseSubIDCounter *int `hl7:"POS=2"`
	DispenseGiveCode CE `hl7:"POS=3"`
	DateTimeDispensed time.Time `hl7:"POS=4"`
	ActualDispenseAmount *int `hl7:"POS=5"`
	ActualDispenseUnits CE `hl7:"POS=6"`
	ActualDosageForm CE `hl7:"POS=7"`
	PrescriptionNumber string `hl7:"POS=8"`
	NumberOfRefillsRemaining *int `hl7:"POS=9"`
	DispenseNotes []string `hl7:"POS=10"`
	DispensingProvider XCN `hl7:"POS=11"`
	SubstitutionStatus string `hl7:"POS=12"`
	TotalDailyDose *int `hl7:"POS=13"`
	DispenseToLocation LA2 `hl7:"POS=14"`
	NeedsHumanReview string `hl7:"POS=15"`
	PharmacyTreatmentSuppliersSpecialDispensingInstructions []CE `hl7:"POS=16"`
	ActualStrength *int `hl7:"POS=17"`
	ActualStrengthUnit CE `hl7:"POS=18"`
	SubstanceLotNumber []string `hl7:"POS=19"`
	SubstanceExpirationDate []time.Time `hl7:"POS=20"`
	SubstanceManufacturerName []CE `hl7:"POS=21"`
	Indication []CE `hl7:"POS=22"`
	DispensePackageSize *int `hl7:"POS=23"`
	DispensePackageSizeUnit CE `hl7:"POS=24"`
	DispensePackageMethod string `hl7:"POS=25"`
}

