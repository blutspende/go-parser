package hl7v271

// SID - Substance Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/SID
type SID struct {
	ApplicationMethodIdentifier CWE `hl7:"POS=2"`
	SubstanceLotNumber string `hl7:"POS=3"`
	SubstanceContainerIdentifier string `hl7:"POS=4"`
	SubstanceManufacturerIdentifier CWE `hl7:"POS=5"`
}

