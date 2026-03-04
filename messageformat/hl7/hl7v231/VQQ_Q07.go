package hl7v231

// VQQ_Q07 - Virtual table query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/VQQ_Q07
type VQQ_Q07 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	VirtualTableQueryRequestSegment VTQ `hl7:"TAG=VTQ"`
	TableRowDefinitionSegment RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

