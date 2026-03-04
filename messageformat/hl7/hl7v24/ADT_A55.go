package hl7v24

// ADT_A55 - Cancel change attending doctor
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ADT_A55
type ADT_A55 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

