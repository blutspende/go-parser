package hl7v231

import "time"

// CSS - Clinical study data schedule segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/CSS
type CSS struct {
	StudyScheduledTimePoint CE `hl7:"POS=2"`
	StudyScheduledPatientTimePoint time.Time `hl7:"POS=3"`
	StudyQualityControlCodes []CE `hl7:"POS=4"`
}

