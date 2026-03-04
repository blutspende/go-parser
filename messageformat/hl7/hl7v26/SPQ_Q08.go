package hl7v26

// SPQ_Q08 - Stored Procedure Request
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/SPQ_Q08
type SPQ_Q08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	StoredProcedureRequestDefinition SPR `hl7:"TAG=SPR"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

