package hl7v27

// SCP - Sterilizer Configuration (anti-microbial Devices)
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/SCP
type SCP struct {
	NumberOfDecontaminationSterilizationDevices *int `hl7:"POS=2"`
	LaborCalculationType CWE `hl7:"POS=3"`
	DateFormat CWE `hl7:"POS=4"`
	DeviceNumber EI `hl7:"POS=5"`
	DeviceName string `hl7:"POS=6"`
	DeviceModelName string `hl7:"POS=7"`
	DeviceType CWE `hl7:"POS=8"`
	LotControl CWE `hl7:"POS=9"`
}

