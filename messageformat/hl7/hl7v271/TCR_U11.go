package hl7v271

// TCR_U11_TEST_CONFIGURATION - Group struct
type TCR_U11_TEST_CONFIGURATION struct {
	Specimen SPM `hl7:"TAG=SPM;ATR=optional"`
	TestCodeConfiguration []TCC `hl7:"TAG=TCC"`
}

// TCR_U11 - Automated equipment test code settings request
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/TCR_U11
type TCR_U11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	Test_configuration []TCR_U11_TEST_CONFIGURATION `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

