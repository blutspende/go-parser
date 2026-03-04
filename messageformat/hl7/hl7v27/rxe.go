package hl7v27

import "time"

// RXE - Pharmacy/treatment Encoded Order
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/RXE
type RXE struct {
	GiveCode CWE `hl7:"POS=3"`
	GiveAmountMinimum *int `hl7:"POS=4"`
	GiveAmountMaximum *int `hl7:"POS=5"`
	GiveUnits CWE `hl7:"POS=6"`
	GiveDosageForm CWE `hl7:"POS=7"`
	ProvidersAdministrationInstructions []CWE `hl7:"POS=8"`
	SubstitutionStatus string `hl7:"POS=10"`
	DispenseAmount *int `hl7:"POS=11"`
	DispenseUnits CWE `hl7:"POS=12"`
	NumberOfRefills *int `hl7:"POS=13"`
	OrderingProvidersDeaNumber []XCN `hl7:"POS=14"`
	PharmacistTreatmentSuppliersVerifierID []XCN `hl7:"POS=15"`
	PrescriptionNumber string `hl7:"POS=16"`
	NumberOfRefillsRemaining *int `hl7:"POS=17"`
	NumberOfRefillsDosesDispensed *int `hl7:"POS=18"`
	DTOfMostRecentRefillOrDoseDispensed time.Time `hl7:"POS=19"`
	TotalDailyDose CQ `hl7:"POS=20"`
	NeedsHumanReview string `hl7:"POS=21"`
	PharmacyTreatmentSuppliersSpecialDispensingInstructions []CWE `hl7:"POS=22"`
	GivePerTimeUnit string `hl7:"POS=23"`
	GiveRateAmount string `hl7:"POS=24"`
	GiveRateUnits CWE `hl7:"POS=25"`
	GiveStrength *int `hl7:"POS=26"`
	GiveStrengthUnits CWE `hl7:"POS=27"`
	GiveIndication []CWE `hl7:"POS=28"`
	DispensePackageSize *int `hl7:"POS=29"`
	DispensePackageSizeUnit CWE `hl7:"POS=30"`
	DispensePackageMethod string `hl7:"POS=31"`
	SupplementaryCode []CWE `hl7:"POS=32"`
	OriginalOrderDateTime time.Time `hl7:"POS=33"`
	GiveDrugStrengthVolume *int `hl7:"POS=34"`
	GiveDrugStrengthVolumeUnits CWE `hl7:"POS=35"`
	ControlledSubstanceSchedule CWE `hl7:"POS=36"`
	FormularyStatus string `hl7:"POS=37"`
	PharmaceuticalSubstanceAlternative []CWE `hl7:"POS=38"`
	PharmacyOfMostRecentFill CWE `hl7:"POS=39"`
	InitialDispenseAmount *int `hl7:"POS=40"`
	DispensingPharmacy CWE `hl7:"POS=41"`
	DispensingPharmacyAddress XAD `hl7:"POS=42"`
	DeliverToPatientLocation PL `hl7:"POS=43"`
	DeliverToAddress XAD `hl7:"POS=44"`
	PharmacyOrderType string `hl7:"POS=45"`
	PharmacyPhoneNumber []XTN `hl7:"POS=46"`
}

