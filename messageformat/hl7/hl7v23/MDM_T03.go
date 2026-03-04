package hl7v23

// MDM_T03 - Document status change notification
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MDM_T03
type MDM_T03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	DocumentNotificationSegment TXA `hl7:"TAG=TXA"`
}

