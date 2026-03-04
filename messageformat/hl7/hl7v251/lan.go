package hl7v251

// LAN - Language Detail
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/LAN
type LAN struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	LanguageCode CE `hl7:"POS=3"`
	LanguageAbilityCode []CE `hl7:"POS=4"`
	LanguageProficiencyCode CE `hl7:"POS=5"`
}

