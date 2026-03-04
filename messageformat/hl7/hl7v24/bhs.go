package hl7v24

import "time"

// BHS - Batch Header Segment
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/BHS
type BHS struct {
	BatchFieldSeparator string `hl7:"POS=2"`
	BatchEncodingCharacters string `hl7:"POS=3"`
	BatchSendingApplication string `hl7:"POS=4"`
	BatchSendingFacility string `hl7:"POS=5"`
	BatchReceivingApplication string `hl7:"POS=6"`
	BatchReceivingFacility string `hl7:"POS=7"`
	BatchCreationDateTime time.Time `hl7:"POS=8"`
	BatchSecurity string `hl7:"POS=9"`
	BatchNameIDType string `hl7:"POS=10"`
	BatchComment string `hl7:"POS=11"`
	BatchControlID string `hl7:"POS=12"`
	ReferenceBatchControlID string `hl7:"POS=13"`
}

