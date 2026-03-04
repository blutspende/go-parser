package hl7v24

import "time"

// EDU - Educational Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/EDU
type EDU struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	AcademicDegree string `hl7:"POS=3"`
	AcademicDegreeProgramDateRange DR `hl7:"POS=4"`
	AcademicDegreeProgramParticipationDateRange DR `hl7:"POS=5"`
	AcademicDegreeGrantedDate time.Time `hl7:"POS=6;ATR=date"`
	School XON `hl7:"POS=7"`
	SchoolTypeCode CE `hl7:"POS=8"`
	SchoolAddress XAD `hl7:"POS=9"`
}

