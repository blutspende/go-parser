package hl7v271

// UAC - User Authentication Credential Segment
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/UAC
type UAC struct {
	UserAuthenticationCredentialTypeCode CWE `hl7:"POS=2"`
	UserAuthenticationCredential ED `hl7:"POS=3"`
}

