package hl7v23

// VTQ - Virtual Table Query Request
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/VTQ
type VTQ struct {
	QueryTag string `hl7:"POS=2"`
	QueryResponseFormatCode string `hl7:"POS=3"`
	VtQueryName CE `hl7:"POS=4"`
	VirtualTableName CE `hl7:"POS=5"`
	SelectionCriteria []QSC `hl7:"POS=6"`
}

