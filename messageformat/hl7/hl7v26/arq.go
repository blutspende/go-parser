package hl7v26

// ARQ - Appointment Request
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/ARQ
type ARQ struct {
	PlacerAppointmentID EI `hl7:"POS=2"`
	FillerAppointmentID EI `hl7:"POS=3"`
	OccurrenceNumber *int `hl7:"POS=4"`
	PlacerGroupNumber EI `hl7:"POS=5"`
	ScheduleID CWE `hl7:"POS=6"`
	RequestEventReason CWE `hl7:"POS=7"`
	AppointmentReason CWE `hl7:"POS=8"`
	AppointmentType CWE `hl7:"POS=9"`
	AppointmentDuration *int `hl7:"POS=10"`
	AppointmentDurationUnits CNE `hl7:"POS=11"`
	RequestedStartDateTimeRange []DR `hl7:"POS=12"`
	PriorityArq string `hl7:"POS=13"`
	RepeatingInterval RI `hl7:"POS=14"`
	RepeatingIntervalDuration string `hl7:"POS=15"`
	PlacerContactPerson []XCN `hl7:"POS=16"`
	PlacerContactPhoneNumber []XTN `hl7:"POS=17"`
	PlacerContactAddress []XAD `hl7:"POS=18"`
	PlacerContactLocation PL `hl7:"POS=19"`
	EnteredByPerson []XCN `hl7:"POS=20"`
	EnteredByPhoneNumber []XTN `hl7:"POS=21"`
	EnteredByLocation PL `hl7:"POS=22"`
	ParentPlacerAppointmentID EI `hl7:"POS=23"`
	ParentFillerAppointmentID EI `hl7:"POS=24"`
	PlacerOrderNumber []EI `hl7:"POS=25"`
	FillerOrderNumber []EI `hl7:"POS=26"`
}

