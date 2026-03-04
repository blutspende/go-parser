package hl7v24

// VQQ_Q07 - Virtual table query
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/VQQ_Q07
type VQQ_Q07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	VirtualTableQueryRequest VTQ `hl7:"TAG=VTQ"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

