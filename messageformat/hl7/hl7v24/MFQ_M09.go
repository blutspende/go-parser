package hl7v24

// MFQ_M09 - Master files query - Test/observation (categorical) master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFQ_M09
type MFQ_M09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

