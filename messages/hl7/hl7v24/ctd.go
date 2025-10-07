package hl7v24

// HL7 v2.4 - CTD - Contact Data
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CTD
type CTD struct {
	ContactRole                     []CE  `hl7:"1" json:"ContactRole,omitempty"`
	ContactName                     []XPN `hl7:"2" json:"Contactname,omitempty"`
	ContactAddress                  []XAD `hl7:"3" json:"ContactAddress,omitempty"`
	ContactLocation                 PL    `hl7:"4" json:"ContactLocation,omitempty"`
	ContactCommunicationInformation []XTN `hl7:"5" json:"contactCommunicationInformation,omitempty"`
	PreferredMethodOfContact        CE    `hl7:"6" json:"PreferredMethodOfContact,omitempty"`
	ContactIdentifier               []PI  `hl7:"7" json:"ContactIdentifier,omitempty"`
}
