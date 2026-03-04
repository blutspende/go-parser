package hl7v24

// SID - Substance Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SID
type SID struct {
	ApplicationMethodIdentifier CE `hl7:"POS=2"`
	SubstanceLotNumber string `hl7:"POS=3"`
	SubstanceContainerIdentifier string `hl7:"POS=4"`
	SubstanceManufacturerIdentifier CE `hl7:"POS=5"`
}

