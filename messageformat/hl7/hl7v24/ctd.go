package hl7v24

// HL7 v2.4 - CTD - Contact Data
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CTD
type CTD struct {
	ContactRole                     []CE  `hl7:"POS=2" json:"contactRole,omitempty"`
	ContactName                     []XPN `hl7:"POS=3" json:"contactName,omitempty"`
	ContactAddress                  []XAD `hl7:"POS=4" json:"contactAddress,omitempty"`
	ContactLocation                 PL    `hl7:"POS=5" json:"contactLocation,omitempty"`
	ContactCommunicationInformation []XTN `hl7:"POS=6" json:"contactCommunicationInformation,omitempty"`
	PreferredMethodOfContact        CE    `hl7:"POS=7" json:"preferredMethodOfContact,omitempty"`
	ContactIdentifier               []PI  `hl7:"POS=8" json:"contactIdentifier,omitempty"`
}
