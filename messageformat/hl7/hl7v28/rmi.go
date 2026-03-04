package hl7v28

import "time"

// RMI - Risk Management Incident
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/RMI
type RMI struct {
	RiskManagementIncidentCode CWE `hl7:"POS=2"`
	DateTimeIncident time.Time `hl7:"POS=3"`
	IncidentTypeCode CWE `hl7:"POS=4"`
}

