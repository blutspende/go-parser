package hl7v27

// ARV - Access Restriction
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/ARV
type ARV struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	AccessRestrictionActionCode CNE `hl7:"POS=3"`
	AccessRestrictionValue CWE `hl7:"POS=4"`
	AccessRestrictionReason []CWE `hl7:"POS=5"`
	SpecialAccessRestrictionInstructions []string `hl7:"POS=6"`
	AccessRestrictionDateRange DR `hl7:"POS=7"`
}

