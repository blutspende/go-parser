package hl7v22

import "time"

// NCK - System Clock
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/NCK
type NCK struct {
	SystemDateTime time.Time `hl7:"POS=2"`
}

