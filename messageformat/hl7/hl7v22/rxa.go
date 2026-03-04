package hl7v22

import "time"

// RXA - Pharmacy Administration
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/RXA
type RXA struct {
	GiveSubIDCounter *int `hl7:"POS=2"`
	AdministrationSubIDCounter *int `hl7:"POS=3"`
	DateTimeStartOfAdministration time.Time `hl7:"POS=4"`
	DateTimeEndOfAdministration time.Time `hl7:"POS=5"`
	AdministeredCode CE `hl7:"POS=6"`
	AdministeredAmount *int `hl7:"POS=7"`
	AdministeredUnits CE `hl7:"POS=8"`
	AdministeredDosageForm CE `hl7:"POS=9"`
	AdministrationNotes []string `hl7:"POS=10"`
	AdministeringProvider CN_PERSON `hl7:"POS=11"`
	AdministeredAtLocation CM_LA1 `hl7:"POS=12"`
	AdministeredPerTimeUnit string `hl7:"POS=13"`
}

