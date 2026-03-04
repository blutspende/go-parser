package hl7v24

// VXX_V02_PATIENT - Group struct
type VXX_V02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
}

// VXX_V02 - Response to vaccination query returning multiple PID matches 
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/VXX_V02
type VXX_V02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient []VXX_V02_PATIENT `hl7:"GROUP"`
}

