package hl7v23

import "time"

// CSS - Clinical Study Data Schedule
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/CSS
type CSS struct {
	StudyScheduledTimePoint CE `hl7:"POS=2"`
	StudyScheduledPatientTimePoint time.Time `hl7:"POS=3"`
	StudyQualityControlCodes []CE `hl7:"POS=4"`
}

