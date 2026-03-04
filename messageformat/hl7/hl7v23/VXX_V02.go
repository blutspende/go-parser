package hl7v23

// VXX_V02_PATIENT - Group struct
type VXX_V02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
}

// VXX_V02 - Response to vaccination query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/VXX_V02
type VXX_V02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient []VXX_V02_PATIENT `hl7:"GROUP"`
}

