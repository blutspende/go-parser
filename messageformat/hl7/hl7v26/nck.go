package hl7v26

import "time"

// NCK - System Clock
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/NCK
type NCK struct {
	SystemDateTime time.Time `hl7:"POS=2"`
}

