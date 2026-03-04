package hl7v271

// SDD - Sterilization Device Data
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/SDD
type SDD struct {
	LotNumber EI `hl7:"POS=2"`
	DeviceNumber EI `hl7:"POS=3"`
	DeviceName string `hl7:"POS=4"`
	DeviceDataState CWE `hl7:"POS=5"`
	LoadStatus CWE `hl7:"POS=6"`
	ControlCode *int `hl7:"POS=7"`
	OperatorName string `hl7:"POS=8"`
}

