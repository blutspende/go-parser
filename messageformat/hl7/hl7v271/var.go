package hl7v271

import "time"

// VAR - Variance
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/VAR
type VAR struct {
	VarianceInstanceID EI `hl7:"POS=2"`
	DocumentedDateTime time.Time `hl7:"POS=3"`
	StatedVarianceDateTime time.Time `hl7:"POS=4"`
	VarianceOriginator []XCN `hl7:"POS=5"`
	VarianceClassification CWE `hl7:"POS=6"`
	VarianceDescription []string `hl7:"POS=7"`
}

