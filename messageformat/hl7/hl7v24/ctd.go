package hl7v24

// CTD - Contact Data
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CTD
type CTD struct {
	ContactRole []CE `hl7:"POS=2"`
	ContactName []XPN `hl7:"POS=3"`
	ContactAddress []XAD `hl7:"POS=4"`
	ContactLocation PL `hl7:"POS=5"`
	ContactCommunicationInformation []XTN `hl7:"POS=6"`
	PreferredMethodOfContact CE `hl7:"POS=7"`
	ContactIdentifiers []PI `hl7:"POS=8"`
}

