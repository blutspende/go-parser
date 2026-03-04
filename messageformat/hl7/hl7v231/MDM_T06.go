package hl7v231

// MDM_T06 - Document addendum notification and content
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MDM_T06
type MDM_T06 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	DocumentNotificationSegment TXA `hl7:"TAG=TXA"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX"`
}

