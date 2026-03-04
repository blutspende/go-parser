package hl7v251

import "time"

// NCK - System Clock
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/NCK
type NCK struct {
	SystemDateTime time.Time `hl7:"POS=2"`
}

