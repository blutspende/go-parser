package hl7v25

import "time"

// RMI - Risk Management Incident
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/RMI
type RMI struct {
	RiskManagementIncidentCode CE `hl7:"POS=2"`
	DateTimeIncident time.Time `hl7:"POS=3"`
	IncidentTypeCode CE `hl7:"POS=4"`
}

