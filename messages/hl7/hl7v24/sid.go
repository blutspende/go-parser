package hl7v24

// HL7 v2.4 - SID - Substance Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SID
type SID struct {
	ApplicationOrMethodIdentifier   CE     `hl7:"1" json:"ApplicationOrMethodIdentifier,omitempty"`
	SubstanceLotNumber              string `hl7:"2" json:"SubstanceLotNumber,omitempty"`
	SubstanceContainerIdentifier    string `hl7:"3" json:"SubstanceContainerIdentifier,omitempty"`
	SubstanceManufacturerIdentifier CE     `hl7:"4" json:"SubstanceManufacturerIdentifier,omitempty"`
}
