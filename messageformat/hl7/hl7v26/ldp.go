package hl7v26

import "time"

// LDP - Location Department
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/LDP
type LDP struct {
	PrimaryKeyValueLdp PL `hl7:"POS=2"`
	LocationDepartment CWE `hl7:"POS=3"`
	LocationService []string `hl7:"POS=4"`
	SpecialtyType []CWE `hl7:"POS=5"`
	ValidPatientClasses []string `hl7:"POS=6"`
	ActiveInactiveFlag string `hl7:"POS=7"`
	ActivationDateLdp time.Time `hl7:"POS=8"`
	InactivationDateLdp time.Time `hl7:"POS=9"`
	InactivatedReason string `hl7:"POS=10"`
	VisitingHours []VH `hl7:"POS=11"`
	ContactPhone XTN `hl7:"POS=12"`
	LocationCostCenter CWE `hl7:"POS=13"`
}

