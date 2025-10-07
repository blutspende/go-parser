package hl7v24

// HL7 v2.4 - TCD - Test Code Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/TCD
type TCD struct {
	UniversalServiceIdentifier            CE     `hl7:"1" json:"UniversalServiceIdentifier,omitempty"`
	AntiDilutionFactor                    SN     `hl7:"2" json:"AntiDilutionFactor,omitempty"`
	RerunDilutionFactor                   SN     `hl7:"3" json:"RerunDilutionFactor,omitempty"`
	PreDilutionFActor                     SN     `hl7:"4" json:"PreDilutionFActor,omitempty"`
	EndogenousContentOfPreDilutionDiluent SN     `hl7:"5" json:"EndogenousContentOfPreDilutionDiluent,omitempty"`
	AutomaticRepeatAlowwed                string `hl7:"6" json:"AutomaticRepeatAlowwed,omitempty"`
	ReflexAllowed                         string `hl7:"7" json:"ReflexAllowed,omitempty"`
	AnalyteRepeatStatus                   CE     `hl7:"8" json:"AnalyteRepeatStatus,omitempty"`
}
