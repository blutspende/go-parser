package hl7v251

// SCH - Scheduling Activity Information
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/SCH
type SCH struct {
	PlacerAppointmentID EI `hl7:"POS=2"`
	FillerAppointmentID EI `hl7:"POS=3"`
	OccurrenceNumber *int `hl7:"POS=4"`
	PlacerGroupNumber EI `hl7:"POS=5"`
	ScheduleID CE `hl7:"POS=6"`
	EventReason CE `hl7:"POS=7"`
	AppointmentReason CE `hl7:"POS=8"`
	AppointmentType CE `hl7:"POS=9"`
	AppointmentDuration *int `hl7:"POS=10"`
	AppointmentDurationUnits CE `hl7:"POS=11"`
	AppointmentTimingQuantity []TQ `hl7:"POS=12"`
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
	FillerStatusCode CE `hl7:"POS=26"`
	PlacerOrderNumber []EI `hl7:"POS=27"`
	FillerOrderNumber []EI `hl7:"POS=28"`
}

