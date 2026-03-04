package hl7v24

// SPQ_Q08 - Stored procedure request
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/SPQ_Q08
type SPQ_Q08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	StoredProcedureRequestDefinition SPR `hl7:"TAG=SPR"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

