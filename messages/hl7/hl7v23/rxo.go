package hl7v23

// RXO - Pharmacy prescription order segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RXO
type RXO struct {
	RequestedGiveCode             CE       `hl7:"POS=2" json:"requestedGiveCode,omitempty"`
	RequestedGiveAmountMinimum    *float32 `hl7:"POS=3" json:"requestedGiveAmountMinimum,omitempty"`
	RequestedGiveAmountMaximum    *float32 `hl7:"POS=4" json:"requestedGiveAmountMaximum,omitempty"`
	RequestedGiveUnits            CE       `hl7:"POS=5" json:"requestedGiveUnits,omitempty"`
	RequestedDosageForm           CE       `hl7:"POS=6" json:"requestedDosageForm,omitempty"`
	ProvidersPharmacyInstructions []CE     `hl7:"POS=7" json:"providersPharmacyInstructions,omitempty"`
	DeliverToLocation             []LA1    `hl7:"POS=8" json:"deliverToLocation,omitempty"`
	AllowSubstitutions            string   `hl7:"POS=9" json:"allowSubstitutions,omitempty"`
	RequestedDispenseCode         CE       `hl7:"POS=10" json:"requestedDispenseCode,omitempty"`
	RequestedDispenseAmount       *float32 `hl7:"POS=11" json:"requestedDispenseAmount,omitempty"`
	RequestedDispenseUnits        CE       `hl7:"POS=12" json:"requestedDispenseUnits,omitempty"`
	NumberOfRefills               *float32 `hl7:"POS=13" json:"numberOfRefills,omitempty"`
	OrderingProvidersDEA          XCN      `hl7:"POS=14" json:"orderingProvidersDEA,omitempty"`
	PharmacistSupplierVerifierID  XCN      `hl7:"POS=15" json:"pharmacistSupplierVerifierID,omitempty"`
	NeedsHumanReview              string   `hl7:"POS=16" json:"needsHumanReview,omitempty"`
	RequestedGivePer              string   `hl7:"POS=17" json:"requestedGivePer,omitempty"`
	RequestedGiveStrength         *float32 `hl7:"POS=18" json:"requestedGiveStrength,omitempty"`
	RequestedGiveStrengthUnits    CE       `hl7:"POS=19" json:"requestedGiveStrengthUnits,omitempty"`
	Indication                    CE       `hl7:"POS=20" json:"indication,omitempty"`
	RequestedGiveRateAmount       string   `hl7:"POS=21" json:"requestedGiveRateAmount,omitempty"`
	RequestedGiveRateUnits        CE       `hl7:"POS=22" json:"requestedGiveRateUnits,omitempty"`
}
