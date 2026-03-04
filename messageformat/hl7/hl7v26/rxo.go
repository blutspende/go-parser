package hl7v26

// RXO - Pharmacy/Treatment Order
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/RXO
type RXO struct {
	RequestedGiveCode CWE `hl7:"POS=2"`
	RequestedGiveAmountMinimum *int `hl7:"POS=3"`
	RequestedGiveAmountMaximum *int `hl7:"POS=4"`
	RequestedGiveUnits CWE `hl7:"POS=5"`
	RequestedDosageForm CWE `hl7:"POS=6"`
	ProvidersPharmacyTreatmentInstructions []CWE `hl7:"POS=7"`
	ProvidersAdministrationInstructions []CWE `hl7:"POS=8"`
	DeliverToLocation LA1 `hl7:"POS=9"`
	AllowSubstitutions string `hl7:"POS=10"`
	RequestedDispenseCode CWE `hl7:"POS=11"`
	RequestedDispenseAmount *int `hl7:"POS=12"`
	RequestedDispenseUnits CWE `hl7:"POS=13"`
	NumberOfRefills *int `hl7:"POS=14"`
	OrderingProvidersDeaNumber []XCN `hl7:"POS=15"`
	PharmacistTreatmentSuppliersVerifierID []XCN `hl7:"POS=16"`
	NeedsHumanReview string `hl7:"POS=17"`
	RequestedGivePerTimeUnit string `hl7:"POS=18"`
	RequestedGiveStrength *int `hl7:"POS=19"`
	RequestedGiveStrengthUnits CWE `hl7:"POS=20"`
	Indication []CWE `hl7:"POS=21"`
	RequestedGiveRateAmount string `hl7:"POS=22"`
	RequestedGiveRateUnits CWE `hl7:"POS=23"`
	TotalDailyDose CQ `hl7:"POS=24"`
	SupplementaryCode []CWE `hl7:"POS=25"`
	RequestedDrugStrengthVolume *int `hl7:"POS=26"`
	RequestedDrugStrengthVolumeUnits CWE `hl7:"POS=27"`
	PharmacyOrderType string `hl7:"POS=28"`
	DispensingInterval *int `hl7:"POS=29"`
	MedicationInstanceIdentifier EI `hl7:"POS=30"`
	SegmentInstanceIdentifier EI `hl7:"POS=31"`
	MoodCode CNE `hl7:"POS=32"`
	DispensingPharmacy CWE `hl7:"POS=33"`
	DispensingPharmacyAddress XAD `hl7:"POS=34"`
	DeliverToPatientLocation PL `hl7:"POS=35"`
	DeliverToAddress XAD `hl7:"POS=36"`
}

