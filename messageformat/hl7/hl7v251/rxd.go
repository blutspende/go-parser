package hl7v251

import "time"

// RXD - Pharmacy/Treatment Dispense
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/RXD
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
	DispensingProvider []XCN `hl7:"POS=11"`
	SubstitutionStatus string `hl7:"POS=12"`
	TotalDailyDose CQ `hl7:"POS=13"`
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
	SupplementaryCode []CE `hl7:"POS=26"`
	InitiatingLocation CE `hl7:"POS=27"`
	PackagingAssemblyLocation CE `hl7:"POS=28"`
	ActualDrugStrengthVolume *int `hl7:"POS=29"`
	ActualDrugStrengthVolumeUnits CWE `hl7:"POS=30"`
	DispenseToPharmacy CWE `hl7:"POS=31"`
	DispenseToPharmacyAddress XAD `hl7:"POS=32"`
	PharmacyOrderType string `hl7:"POS=33"`
	DispenseType CWE `hl7:"POS=34"`
}

