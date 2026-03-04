package hl7v26

// UAC - User Authentication Credential
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/UAC
type UAC struct {
	UserAuthenticationCredentialTypeCode CWE `hl7:"POS=2"`
	UserAuthenticationCredential ED `hl7:"POS=3"`
}

