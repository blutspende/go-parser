package hl7v231

// OSQ_Q06 - Query for order status
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/OSQ_Q06
type OSQ_Q06 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

