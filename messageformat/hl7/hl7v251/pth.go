package hl7v251

import "time"

// PTH - Pathway
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/PTH
type PTH struct {
	ActionCode string `hl7:"POS=2"`
	PathwayID CE `hl7:"POS=3"`
	PathwayInstanceID EI `hl7:"POS=4"`
	PathwayEstablishedDateTime time.Time `hl7:"POS=5"`
	PathwayLifeCycleStatus CE `hl7:"POS=6"`
	ChangePathwayLifeCycleStatusDateTime time.Time `hl7:"POS=7"`
}

