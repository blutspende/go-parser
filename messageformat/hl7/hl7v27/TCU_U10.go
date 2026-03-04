package hl7v27

// TCU_U10_TEST_CONFIGURATION - Group struct
type TCU_U10_TEST_CONFIGURATION struct {
	Specimen SPM `hl7:"TAG=SPM;ATR=optional"`
	TestCodeConfiguration []TCC `hl7:"TAG=TCC"`
}

// TCU_U10 - Automated equipment test code settings update
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/TCU_U10
type TCU_U10 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	Test_configuration []TCU_U10_TEST_CONFIGURATION `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

