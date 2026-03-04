package hl7v251

import "time"

// CSS - Clinical Study Data Schedule Segment
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/CSS
type CSS struct {
	StudyScheduledTimePoint CE `hl7:"POS=2"`
	StudyScheduledPatientTimePoint time.Time `hl7:"POS=3"`
	StudyQualityControlCodes []CE `hl7:"POS=4"`
}

