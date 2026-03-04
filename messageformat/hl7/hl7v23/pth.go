package hl7v23

import "time"

// PTH - Pathway
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PTH
type PTH struct {
	ActionCode string `hl7:"POS=2"`
	PathwayID CE `hl7:"POS=3"`
	PathwayInstanceID EI `hl7:"POS=4"`
	PathwayEstablishedDateTime time.Time `hl7:"POS=5"`
	PathwayLifecycleStatus CE `hl7:"POS=6"`
	ChangePathwayLifecycleStatusDateTime time.Time `hl7:"POS=7"`
}

