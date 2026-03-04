package hl7v24

// MFQ_M01 - Master files query - Master file not otherwise specified
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFQ_M01
type MFQ_M01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

