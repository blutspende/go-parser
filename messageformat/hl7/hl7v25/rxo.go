package hl7v25

// RXO - Pharmacy/Treatment Order
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/RXO
type RXO struct {
	RequestedGiveCode CE `hl7:"POS=2"`
	RequestedGiveAmountMinimum *int `hl7:"POS=3"`
	RequestedGiveAmountMaximum *int `hl7:"POS=4"`
	RequestedGiveUnits CE `hl7:"POS=5"`
	RequestedDosageForm CE `hl7:"POS=6"`
	ProvidersPharmacyTreatmentInstructions []CE `hl7:"POS=7"`
	ProvidersAdministrationInstructions []CE `hl7:"POS=8"`
	DeliverToLocation LA1 `hl7:"POS=9"`
	AllowSubstitutions string `hl7:"POS=10"`
	RequestedDispenseCode CE `hl7:"POS=11"`
	RequestedDispenseAmount *int `hl7:"POS=12"`
	RequestedDispenseUnits CE `hl7:"POS=13"`
	NumberOfRefills *int `hl7:"POS=14"`
	OrderingProvidersDeaNumber []XCN `hl7:"POS=15"`
	PharmacistTreatmentSuppliersVerifierID []XCN `hl7:"POS=16"`
	NeedsHumanReview string `hl7:"POS=17"`
	RequestedGivePerTimeUnit string `hl7:"POS=18"`
	RequestedGiveStrength *int `hl7:"POS=19"`
	RequestedGiveStrengthUnits CE `hl7:"POS=20"`
	Indication []CE `hl7:"POS=21"`
	RequestedGiveRateAmount string `hl7:"POS=22"`
	RequestedGiveRateUnits CE `hl7:"POS=23"`
	TotalDailyDose CQ `hl7:"POS=24"`
	SupplementaryCode []CE `hl7:"POS=25"`
	RequestedDrugStrengthVolume *int `hl7:"POS=26"`
	RequestedDrugStrengthVolumeUnits CWE `hl7:"POS=27"`
	PharmacyOrderType string `hl7:"POS=28"`
	DispensingInterval *int `hl7:"POS=29"`
}

