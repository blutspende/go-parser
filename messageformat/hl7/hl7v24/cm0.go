package hl7v24

import "time"

// CM0 - Clinical Study Master
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CM0
type CM0 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SponsorStudyID EI `hl7:"POS=3"`
	AlternateStudyID []EI `hl7:"POS=4"`
	TitleOfStudy string `hl7:"POS=5"`
	ChairmanOfStudy []XCN `hl7:"POS=6"`
	LastIrbApprovalDate time.Time `hl7:"POS=7;ATR=date"`
	TotalAccrualToDate *int `hl7:"POS=8"`
	LastAccrualDate time.Time `hl7:"POS=9;ATR=date"`
	ContactForStudy []XCN `hl7:"POS=10"`
	ContactsTelephoneNumber XTN `hl7:"POS=11"`
	ContactsAddress []XAD `hl7:"POS=12"`
}

