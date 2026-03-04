package hl7v22

// RXO - Pharmacy Prescription Order
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/RXO
type RXO struct {
	RequestedGiveCode CE `hl7:"POS=2"`
	RequestedGiveAmountMinimum *int `hl7:"POS=3"`
	RequestedGiveAmountMaximum *int `hl7:"POS=4"`
	RequestedGiveUnits CE `hl7:"POS=5"`
	RequestedDosageForm CE `hl7:"POS=6"`
	ProvidersPharmacyInstructions []CE `hl7:"POS=7"`
	ProvidersAdministrationInstructions []CE `hl7:"POS=8"`
	DeliverToLocation CM_LA1 `hl7:"POS=9"`
	AllowSubstitutions string `hl7:"POS=10"`
	RequestedDispenseCode CE `hl7:"POS=11"`
	RequestedDispenseAmount *int `hl7:"POS=12"`
	RequestedDispenseUnits CE `hl7:"POS=13"`
	NumberOfRefills *int `hl7:"POS=14"`
	OrderingProvidersDeaNumber CN_PERSON `hl7:"POS=15"`
	PharmacistVerifierID CN_PERSON `hl7:"POS=16"`
	NeedsHumanReview string `hl7:"POS=17"`
	RequestedGivePerTimeUnit string `hl7:"POS=18"`
}

