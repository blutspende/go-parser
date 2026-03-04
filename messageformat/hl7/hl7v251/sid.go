package hl7v251

// SID - Substance Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/SID
type SID struct {
	ApplicationMethodIdentifier CE `hl7:"POS=2"`
	SubstanceLotNumber string `hl7:"POS=3"`
	SubstanceContainerIdentifier string `hl7:"POS=4"`
	SubstanceManufacturerIdentifier CE `hl7:"POS=5"`
}

