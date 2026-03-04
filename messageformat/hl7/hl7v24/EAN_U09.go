package hl7v24

// EAN_U09_NOTIFICATION - Group struct
type EAN_U09_NOTIFICATION struct {
	NotificationDetail NDS `hl7:"TAG=NDS"`
	NotesAndComments NTE `hl7:"TAG=NTE;ATR=optional"`
}

// EAN_U09 - Automated equipment notification
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/EAN_U09
type EAN_U09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EquipmentDetail EQU `hl7:"TAG=EQU"`
	Notification []EAN_U09_NOTIFICATION `hl7:"GROUP"`
	Role ROL `hl7:"TAG=ROL;ATR=optional"`
}

