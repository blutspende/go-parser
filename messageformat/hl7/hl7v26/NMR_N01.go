package hl7v26

// NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT - Group struct
type NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT struct {
	SystemClock NCK `hl7:"TAG=NCK;ATR=optional"`
	NotesAndComments1 []NTE `hl7:"TAG=NTE;ATR=optional"`
	ApplicationControlLevelStatistics NST `hl7:"TAG=NST;ATR=optional"`
	NotesAndComments2 []NTE `hl7:"TAG=NTE;ATR=optional"`
	ApplicationStatusChange NSC `hl7:"TAG=NSC;ATR=optional"`
	NotesAndComments3 []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMR_N01 - Application Management Response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/NMR_N01
type NMR_N01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD;ATR=optional"`
	ClockAndStatsWithNotesAlt []NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT `hl7:"GROUP"`
}

