package hl7v23

// NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT - Group struct
type NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT struct {
	SystemClock NCK `hl7:"TAG=NCK;ATR=optional"`
	NotesAndCommentsSegment1 []NTE `hl7:"TAG=NTE;ATR=optional"`
	Statistics NST `hl7:"TAG=NST;ATR=optional"`
	NotesAndCommentsSegment2 []NTE `hl7:"TAG=NTE;ATR=optional"`
	StatusChange NSC `hl7:"TAG=NSC;ATR=optional"`
	NotesAndCommentsSegment3 []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMR_N01 - Network management response message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/NMR_N01
type NMR_N01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD;ATR=optional"`
	ClockAndStatsWithNotesAlt []NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT `hl7:"GROUP"`
}

