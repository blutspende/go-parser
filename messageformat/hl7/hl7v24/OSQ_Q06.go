package hl7v24

// OSQ_Q06 - Query for order status
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/OSQ_Q06
type OSQ_Q06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

