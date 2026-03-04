package hl7v24

import "time"

// NCK - System clock
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NCK
type NCK struct {
	SystemDateTime time.Time `hl7:"POS=2"`
}

