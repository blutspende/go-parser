package hl7v231

// SPQ_Q08 - Stored procedure request
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/SPQ_Q08
type SPQ_Q08 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	StoredProcedureRequestDefinitionSegment SPR `hl7:"TAG=SPR"`
	TableRowDefinitionSegment RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

