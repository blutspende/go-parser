package hl7v27

import "time"

// NCK - System Clock
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/NCK
type NCK struct {
	SystemDateTime time.Time `hl7:"POS=2"`
}

