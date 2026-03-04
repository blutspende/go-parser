package hl7v25

// BLC - Blood Code
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/BLC
type BLC struct {
	BloodProductCode CE `hl7:"POS=2"`
	BloodAmount CQ `hl7:"POS=3"`
}

