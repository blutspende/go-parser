package hl7v231

import "time"

// RXA - Pharmacy/treatment administration segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/RXA
type RXA struct {
	GiveSubIDCounter *int `hl7:"POS=2"`
	AdministrationSubIDCounter *int `hl7:"POS=3"`
	DateTimeStartOfAdministration time.Time `hl7:"POS=4"`
	DateTimeEndOfAdministration time.Time `hl7:"POS=5"`
	AdministeredCode CE `hl7:"POS=6"`
	AdministeredAmount *int `hl7:"POS=7"`
	AdministeredUnits CE `hl7:"POS=8"`
	AdministeredDosageForm CE `hl7:"POS=9"`
	AdministrationNotes []CE `hl7:"POS=10"`
	AdministeringProvider []XCN `hl7:"POS=11"`
	AdministeredAtLocation LA2 `hl7:"POS=12"`
	AdministeredPer string `hl7:"POS=13"`
	AdministeredStrength *int `hl7:"POS=14"`
	AdministeredStrengthUnits CE `hl7:"POS=15"`
	SubstanceLotNumber []string `hl7:"POS=16"`
	SubstanceExpirationDate []time.Time `hl7:"POS=17"`
	SubstanceManufacturerName []CE `hl7:"POS=18"`
	SubstanceRefusalReason []CE `hl7:"POS=19"`
	Indication []CE `hl7:"POS=20"`
	CompletionStatus string `hl7:"POS=21"`
	ActionCodeRxa string `hl7:"POS=22"`
	SystemEntryDateTime time.Time `hl7:"POS=23"`
}

