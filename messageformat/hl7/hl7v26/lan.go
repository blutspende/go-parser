package hl7v26

// LAN - Language Detail
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/LAN
type LAN struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	LanguageCode CWE `hl7:"POS=3"`
	LanguageAbilityCode []CWE `hl7:"POS=4"`
	LanguageProficiencyCode CWE `hl7:"POS=5"`
}

