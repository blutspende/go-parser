package hl7v24

// HL7 v2.4 - SID - Substance Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SID
type SID struct {
	ApplicationOrMethodIdentifier   CE     `hl7:"POS=2" json:"ApplicationOrMethodIdentifier,omitempty"`
	SubstanceLotNumber              string `hl7:"POS=3" json:"SubstanceLotNumber,omitempty"`
	SubstanceContainerIdentifier    string `hl7:"POS=4" json:"SubstanceContainerIdentifier,omitempty"`
	SubstanceManufacturerIdentifier CE     `hl7:"POS=5" json:"SubstanceManufacturerIdentifier,omitempty"`
}
