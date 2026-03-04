package hl7v231

// QRY_Q02 - Query sent for deferred response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/QRY_Q02
type QRY_Q02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

