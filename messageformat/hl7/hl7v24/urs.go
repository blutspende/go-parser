package hl7v24

import "time"

// URS - Unsolicited Selection
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/URS
type URS struct {
	RUWhereSubjectDefinition []string `hl7:"POS=2"`
	RUWhenDataStartDateTime time.Time `hl7:"POS=3"`
	RUWhenDataEndDateTime time.Time `hl7:"POS=4"`
	RUWhatUserQualifier []string `hl7:"POS=5"`
	RUOtherResultsSubjectDefinition []string `hl7:"POS=6"`
	RUWhichDateTimeQualifier []string `hl7:"POS=7"`
	RUWhichDateTimeStatusQualifier []string `hl7:"POS=8"`
	RUDateTimeSelectionQualifier []string `hl7:"POS=9"`
	RUQuantityTimingQualifier TQ `hl7:"POS=10"`
}

