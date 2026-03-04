package hl7v25

import "time"

// VAR - Variance
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/VAR
type VAR struct {
	VarianceInstanceID EI `hl7:"POS=2"`
	DocumentedDateTime time.Time `hl7:"POS=3"`
	StatedVarianceDateTime time.Time `hl7:"POS=4"`
	VarianceOriginator []XCN `hl7:"POS=5"`
	VarianceClassification CE `hl7:"POS=6"`
	VarianceDescription []string `hl7:"POS=7"`
}

