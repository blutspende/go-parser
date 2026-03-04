package hl7v23

// OSQ_Q06 - Query for order status
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/OSQ_Q06
type OSQ_Q06 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

