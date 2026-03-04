package hl7v27

// SCH - Scheduling Activity Information
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/SCH
type SCH struct {
	PlacerAppointmentID EI `hl7:"POS=2"`
	FillerAppointmentID EI `hl7:"POS=3"`
	OccurrenceNumber *int `hl7:"POS=4"`
	PlacerGroupNumber EI `hl7:"POS=5"`
	ScheduleID CWE `hl7:"POS=6"`
	EventReason CWE `hl7:"POS=7"`
	AppointmentReason CWE `hl7:"POS=8"`
	AppointmentType CWE `hl7:"POS=9"`
	AppointmentDurationUnits CNE `hl7:"POS=11"`
	PlacerContactPerson []XCN `hl7:"POS=13"`
	PlacerContactPhoneNumber XTN `hl7:"POS=14"`
	PlacerContactAddress []XAD `hl7:"POS=15"`
	PlacerContactLocation PL `hl7:"POS=16"`
	FillerContactPerson []XCN `hl7:"POS=17"`
	FillerContactPhoneNumber XTN `hl7:"POS=18"`
	FillerContactAddress []XAD `hl7:"POS=19"`
	FillerContactLocation PL `hl7:"POS=20"`
	EnteredByPerson []XCN `hl7:"POS=21"`
	EnteredByPhoneNumber []XTN `hl7:"POS=22"`
	EnteredByLocation PL `hl7:"POS=23"`
	ParentPlacerAppointmentID EI `hl7:"POS=24"`
	ParentFillerAppointmentID EI `hl7:"POS=25"`
	FillerStatusCode CWE `hl7:"POS=26"`
	PlacerOrderNumber []EI `hl7:"POS=27"`
	FillerOrderNumber []EI `hl7:"POS=28"`
}

