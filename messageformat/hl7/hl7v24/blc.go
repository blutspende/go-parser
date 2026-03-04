package hl7v24

// BLC - Blood Code
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/BLC
type BLC struct {
	BloodProductCode CE `hl7:"POS=2"`
	BloodAmount CQ `hl7:"POS=3"`
}

