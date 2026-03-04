package hl7v22

// NMR_N02_RESPONSE - Group struct
type NMR_N02_RESPONSE struct {
	SystemClock NCK `hl7:"TAG=NCK;ATR=optional"`
	NotesAndComments1 []NTE `hl7:"TAG=NTE;ATR=optional"`
	Statistics NST `hl7:"TAG=NST;ATR=optional"`
	NotesAndComments2 []NTE `hl7:"TAG=NTE;ATR=optional"`
	StatusChange NSC `hl7:"TAG=NSC;ATR=optional"`
	NotesAndComments3 []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMR_N02 - Network Management Response
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/NMR_N02
type NMR_N02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryDefinition QRD `hl7:"TAG=QRD;ATR=optional"`
	Response []NMR_N02_RESPONSE `hl7:"GROUP"`
}

