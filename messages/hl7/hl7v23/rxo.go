package hl7v23

// RXO - Pharmacy prescription order segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RXO
type RXO struct {
	RequestedGiveCode             CE      `hl7:"2" json:"requestedGiveCode,omitempty"`
	RequestedGiveAmountMinimum    float32 `hl7:"3" json:"requestedGiveAmountMinimum,omitempty"`
	RequestedGiveAmountMaximum    float32 `hl7:"4" json:"requestedGiveAmountMaximum,omitempty"`
	RequestedGiveUnits            CE      `hl7:"5" json:"requestedGiveUnits,omitempty"`
	RequestedDosageForm           CE      `hl7:"6" json:"requestedDosageForm,omitempty"`
	ProvidersPharmacyInstructions []CE    `hl7:"7" json:"providersPharmacyInstructions,omitempty"`
	DeliverToLocation             []LA1   `hl7:"8" json:"deliverToLocation,omitempty"`
	AllowSubstitutions            string  `hl7:"9" json:"allowSubstitutions,omitempty"`
	RequestedDispenseCode         CE      `hl7:"10" json:"requestedDispenseCode,omitempty"`
	RequestedDispenseAmount       float32 `hl7:"11" json:"requestedDispenseAmount,omitempty"`
	RequestedDispenseUnits        CE      `hl7:"12" json:"requestedDispenseUnits,omitempty"`
	NumberOfRefills               float32 `hl7:"13" json:"numberOfRefills,omitempty"`
	OrderingProvidersDEA          XCN     `hl7:"14" json:"orderingProvidersDEA,omitempty"`
	PharmacistSupplierVerifierID  XCN     `hl7:"15" json:"pharmacistSupplierVerifierID,omitempty"`
	NeedsHumanReview              string  `hl7:"16" json:"needsHumanReview,omitempty"`
	RequestedGivePer              string  `hl7:"17" json:"requestedGivePer,omitempty"`
	RequestedGiveStrength         float32 `hl7:"18" json:"requestedGiveStrength,omitempty"`
	RequestedGiveStrengthUnits    CE      `hl7:"19" json:"requestedGiveStrengthUnits,omitempty"`
	Indication                    CE      `hl7:"20" json:"indication,omitempty"`
	RequestedGiveRateAmount       string  `hl7:"21" json:"requestedGiveRateAmount,omitempty"`
	RequestedGiveRateUnits        CE      `hl7:"22" json:"requestedGiveRateUnits,omitempty"`
}
