package hl7v23

// SPQ_Q01 - Stored procedure request
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/SPQ_Q01
type SPQ_Q01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	StoredProcedureRequestDefinition SPR `hl7:"TAG=SPR"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

