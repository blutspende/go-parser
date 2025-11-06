package hl7v24

// HL7 v2.4 - TCD - Test Code Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/TCD
type TCD struct {
	UniversalServiceIdentifier            CE     `hl7:"POS=2" json:"UniversalServiceIdentifier,omitempty"`
	AntiDilutionFactor                    SN     `hl7:"POS=3" json:"AntiDilutionFactor,omitempty"`
	RerunDilutionFactor                   SN     `hl7:"POS=4" json:"RerunDilutionFactor,omitempty"`
	PreDilutionFActor                     SN     `hl7:"POS=5" json:"PreDilutionFActor,omitempty"`
	EndogenousContentOfPreDilutionDiluent SN     `hl7:"POS=6" json:"EndogenousContentOfPreDilutionDiluent,omitempty"`
	AutomaticRepeatAlowwed                string `hl7:"POS=7" json:"AutomaticRepeatAlowwed,omitempty"`
	ReflexAllowed                         string `hl7:"POS=8" json:"ReflexAllowed,omitempty"`
	AnalyteRepeatStatus                   CE     `hl7:"POS=9" json:"AnalyteRepeatStatus,omitempty"`
}
