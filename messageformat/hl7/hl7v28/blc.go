package hl7v28

// BLC - Blood Code
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/BLC
type BLC struct {
	BloodProductCode CWE `hl7:"POS=2"`
	BloodAmount CQ `hl7:"POS=3"`
}

