package hl7v26

// CM2 - Clinical Study Schedule Master
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/CM2
type CM2 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ScheduledTimePoint CWE `hl7:"POS=3"`
	DescriptionOfTimePoint string `hl7:"POS=4"`
	EventsScheduledThisTimePoint []CWE `hl7:"POS=5"`
}

