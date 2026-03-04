package hl7v231

// VXX_V02_PATIENT - Group struct
type VXX_V02_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
}

// VXX_V02 - Response to vaccination query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/VXX_V02
type VXX_V02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient []VXX_V02_PATIENT `hl7:"GROUP"`
}

