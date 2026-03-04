package hl7v26

// EAN_U09_NOTIFICATION - Group struct
type EAN_U09_NOTIFICATION struct {
	NotificationDetail NDS `hl7:"TAG=NDS"`
	NotesAndComments NTE `hl7:"TAG=NTE;ATR=optional"`
}

// EAN_U09 - Automated Equipment Notification
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/EAN_U09
type EAN_U09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	Notification []EAN_U09_NOTIFICATION `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

