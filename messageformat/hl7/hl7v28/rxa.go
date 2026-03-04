package hl7v28

import "time"

// RXA - Pharmacy/treatment Administration
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/RXA
type RXA struct {
	GiveSubIDCounter *int `hl7:"POS=2"`
	AdministrationSubIDCounter *int `hl7:"POS=3"`
	DateTimeStartOfAdministration time.Time `hl7:"POS=4"`
	DateTimeEndOfAdministration time.Time `hl7:"POS=5"`
	AdministeredCode CWE `hl7:"POS=6"`
	AdministeredAmount *int `hl7:"POS=7"`
	AdministeredUnits CWE `hl7:"POS=8"`
	AdministeredDosageForm CWE `hl7:"POS=9"`
	AdministrationNotes []CWE `hl7:"POS=10"`
	AdministeringProvider []XCN `hl7:"POS=11"`
	AdministeredPerTimeUnit string `hl7:"POS=13"`
	AdministeredStrength *int `hl7:"POS=14"`
	AdministeredStrengthUnits CWE `hl7:"POS=15"`
	SubstanceLotNumber []string `hl7:"POS=16"`
	SubstanceExpirationDate []time.Time `hl7:"POS=17"`
	SubstanceManufacturerName []CWE `hl7:"POS=18"`
	SubstanceTreatmentRefusalReason []CWE `hl7:"POS=19"`
	Indication []CWE `hl7:"POS=20"`
	CompletionStatus string `hl7:"POS=21"`
	ActionCodeRxa string `hl7:"POS=22"`
	SystemEntryDateTime time.Time `hl7:"POS=23"`
	AdministeredDrugStrengthVolume *int `hl7:"POS=24"`
	AdministeredDrugStrengthVolumeUnits CWE `hl7:"POS=25"`
	AdministeredBarcodeIdentifier CWE `hl7:"POS=26"`
	PharmacyOrderType string `hl7:"POS=27"`
	AdministerAt PL `hl7:"POS=28"`
	AdministeredAtAddress XAD `hl7:"POS=29"`
	AdministeredTagIdentifier []EI `hl7:"POS=30"`
}

