package hl7v24

// MFQ_M05 - Master files query - Patient location master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFQ_M05
type MFQ_M05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

