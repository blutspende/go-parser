package hl7v23

// VQQ_Q01 - Virtual table query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/VQQ_Q01
type VQQ_Q01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	VirtualTableQueryRequest VTQ `hl7:"TAG=VTQ"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

