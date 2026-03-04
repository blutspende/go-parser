package hl7v251

// VQQ_Q07 - Virtual Table Query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/VQQ_Q07
type VQQ_Q07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	VirtualTableQueryRequest VTQ `hl7:"TAG=VTQ"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

