package hl7v26

// STZ - Sterilization Parameter
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/STZ
type STZ struct {
	SterilizationType CWE `hl7:"POS=2"`
	SterilizationCycle CWE `hl7:"POS=3"`
	MaintenanceCycle CWE `hl7:"POS=4"`
	MaintenanceType CWE `hl7:"POS=5"`
}

