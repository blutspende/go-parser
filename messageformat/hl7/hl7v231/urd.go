package hl7v231

import "time"

// URD - Results/update definition segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/URD
type URD struct {
	RUDateTime time.Time `hl7:"POS=2"`
	ReportPriority string `hl7:"POS=3"`
	RUWhoSubjectDefinition []XCN `hl7:"POS=4"`
	RUWhatSubjectDefinition []CE `hl7:"POS=5"`
	RUWhatDepartmentCode []CE `hl7:"POS=6"`
	RUDisplayPrintLocations []string `hl7:"POS=7"`
	RUResultsLevel string `hl7:"POS=8"`
}

