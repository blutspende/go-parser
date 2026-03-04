package hl7v26

// BLC - Blood Code
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/BLC
type BLC struct {
	BloodProductCode CWE `hl7:"POS=2"`
	BloodAmount CQ `hl7:"POS=3"`
}

