package hl7v271

import "time"

// RXG - Pharmacy/treatment Give
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/RXG
type RXG struct {
	GiveSubIDCounter *int `hl7:"POS=2"`
	DispenseSubIDCounter *int `hl7:"POS=3"`
	GiveCode CWE `hl7:"POS=5"`
	GiveAmountMinimum *int `hl7:"POS=6"`
	GiveAmountMaximum *int `hl7:"POS=7"`
	GiveUnits CWE `hl7:"POS=8"`
	GiveDosageForm CWE `hl7:"POS=9"`
	AdministrationNotes []CWE `hl7:"POS=10"`
	SubstitutionStatus string `hl7:"POS=11"`
	DispenseToLocation LA2 `hl7:"POS=12"`
	NeedsHumanReview string `hl7:"POS=13"`
	PharmacyTreatmentSuppliersSpecialAdministrationInstructions []CWE `hl7:"POS=14"`
	GivePerTimeUnit string `hl7:"POS=15"`
	GiveRateAmount string `hl7:"POS=16"`
	GiveRateUnits CWE `hl7:"POS=17"`
	GiveStrength *int `hl7:"POS=18"`
	GiveStrengthUnits CWE `hl7:"POS=19"`
	SubstanceLotNumber []string `hl7:"POS=20"`
	SubstanceExpirationDate []time.Time `hl7:"POS=21"`
	SubstanceManufacturerName []CWE `hl7:"POS=22"`
	Indication []CWE `hl7:"POS=23"`
	GiveDrugStrengthVolume *int `hl7:"POS=24"`
	GiveDrugStrengthVolumeUnits CWE `hl7:"POS=25"`
	GiveBarcodeIdentifier CWE `hl7:"POS=26"`
	PharmacyOrderType string `hl7:"POS=27"`
	DispenseToPharmacy CWE `hl7:"POS=28"`
	DispenseToPharmacyAddress XAD `hl7:"POS=29"`
	DeliverToPatientLocation PL `hl7:"POS=30"`
	DeliverToAddress XAD `hl7:"POS=31"`
}

