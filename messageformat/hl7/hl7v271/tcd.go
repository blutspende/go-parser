package hl7v271

// TCD - Test Code Detail
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/TCD
type TCD struct {
	UniversalServiceIdentifier CWE `hl7:"POS=2"`
	AutoDilutionFactor SN `hl7:"POS=3"`
	RerunDilutionFactor SN `hl7:"POS=4"`
	PreDilutionFactor SN `hl7:"POS=5"`
	EndogenousContentOfPreDilutionDiluent SN `hl7:"POS=6"`
	AutomaticRepeatAllowed string `hl7:"POS=7"`
	ReflexAllowed string `hl7:"POS=8"`
	AnalyteRepeatStatus CWE `hl7:"POS=9"`
}

