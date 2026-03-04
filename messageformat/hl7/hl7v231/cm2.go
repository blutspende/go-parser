package hl7v231

// CM2 - Clinical study schedule master segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/CM2
type CM2 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ScheduledTimePoint CE `hl7:"POS=3"`
	DescriptionOfTimePoint string `hl7:"POS=4"`
	EventsScheduledThisTimePoint []CE `hl7:"POS=5"`
}

