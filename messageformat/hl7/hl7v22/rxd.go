package hl7v22

import "time"

// RXD - Pharmacy Dispense
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/RXD
type RXD struct {
	DispenseSubIDCounter *int `hl7:"POS=2"`
	DispenseGiveCode CE `hl7:"POS=3"`
	DateTimeDispensed time.Time `hl7:"POS=4"`
	ActualDispenseAmount *int `hl7:"POS=5"`
	ActualDispenseUnits CE `hl7:"POS=6"`
	ActualDosageForm CE `hl7:"POS=7"`
	PrescriptionNumber *int `hl7:"POS=8"`
	NumberOfRefillsRemaining *int `hl7:"POS=9"`
	DispenseNotes []string `hl7:"POS=10"`
	DispensingProvider CN_PERSON `hl7:"POS=11"`
	SubstitutionStatus string `hl7:"POS=12"`
	TotalDailyDose *int `hl7:"POS=13"`
	DispenseToLocation CM_LA1 `hl7:"POS=14"`
	NeedsHumanReview string `hl7:"POS=15"`
	PharmacySpecialDispensingInstructions []CE `hl7:"POS=16"`
}

