package hl7v251

// TCC - Test Code Configuration
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/TCC
type TCC struct {
	UniversalServiceIdentifier CE `hl7:"POS=2"`
	TestApplicationIdentifier EI `hl7:"POS=3"`
	SpecimenSource SPS `hl7:"POS=4"`
	AutoDilutionFactorDefault SN `hl7:"POS=5"`
	RerunDilutionFactorDefault SN `hl7:"POS=6"`
	PreDilutionFactorDefault SN `hl7:"POS=7"`
	EndogenousContentOfPreDilutionDiluent SN `hl7:"POS=8"`
	InventoryLimitsWarningLevel *int `hl7:"POS=9"`
	AutomaticRerunAllowed string `hl7:"POS=10"`
	AutomaticRepeatAllowed string `hl7:"POS=11"`
	AutomaticReflexAllowed string `hl7:"POS=12"`
	EquipmentDynamicRange SN `hl7:"POS=13"`
	Units CE `hl7:"POS=14"`
	ProcessingType CE `hl7:"POS=15"`
}

