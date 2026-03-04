package hl7v26

import "time"

// URD - Results/update Definition
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/URD
type URD struct {
	RUDateTime time.Time `hl7:"POS=2"`
	ReportPriority string `hl7:"POS=3"`
	RUWhoSubjectDefinition []XCN `hl7:"POS=4"`
	RUWhatSubjectDefinition []CWE `hl7:"POS=5"`
	RUWhatDepartmentCode []CWE `hl7:"POS=6"`
	RUDisplayPrintLocations []string `hl7:"POS=7"`
	RUResultsLevel string `hl7:"POS=8"`
}

