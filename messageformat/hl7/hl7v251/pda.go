package hl7v251

import "time"

// PDA - Patient Death and Autopsy
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/PDA
type PDA struct {
	DeathCauseCode []CE `hl7:"POS=2"`
	DeathLocation PL `hl7:"POS=3"`
	DeathCertifiedIndicator string `hl7:"POS=4"`
	DeathCertificateSignedDateTime time.Time `hl7:"POS=5"`
	DeathCertifiedBy XCN `hl7:"POS=6"`
	AutopsyIndicator string `hl7:"POS=7"`
	AutopsyStartAndEndDateTime DR `hl7:"POS=8"`
	AutopsyPerformedBy XCN `hl7:"POS=9"`
	CoronerIndicator string `hl7:"POS=10"`
}

