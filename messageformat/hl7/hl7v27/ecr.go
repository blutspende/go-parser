package hl7v27

import "time"

// ECR - Equipment Command Response
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/ECR
type ECR struct {
	CommandResponse CWE `hl7:"POS=2"`
	DateTimeCompleted time.Time `hl7:"POS=3"`
	CommandResponseParameters []string `hl7:"POS=4"`
}

