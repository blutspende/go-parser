package hl7v271

// CCQ_I19_PROVIDER_CONTACT - Group struct
type CCQ_I19_PROVIDER_CONTACT struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// CCQ_I19 - Collaborative Care Query/Collaborative Care Query Update
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/CCQ_I19
type CCQ_I19 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	ReferralInformation RF1 `hl7:"TAG=RF1"`
	Provider_contact []CCQ_I19_PROVIDER_CONTACT `hl7:"GROUP;ATR=optional"`
	ClinicalRelationshipSegment []REL `hl7:"TAG=REL;ATR=optional"`
}

