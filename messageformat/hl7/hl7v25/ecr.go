package hl7v25

import "time"

// ECR - Equipment Command Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/ECR
type ECR struct {
	CommandResponse CE `hl7:"POS=2"`
	DateTimeCompleted time.Time `hl7:"POS=3"`
	CommandResponseParameters []string `hl7:"POS=4"`
}

