package hl7v27

import "time"

// CSS - Clinical Study Data Schedule Segment
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/CSS
type CSS struct {
	StudyScheduledTimePoint CWE `hl7:"POS=2"`
	StudyScheduledPatientTimePoint time.Time `hl7:"POS=3"`
	StudyQualityControlCodes []CWE `hl7:"POS=4"`
}

