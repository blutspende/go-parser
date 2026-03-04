package hl7v231

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES - Group struct
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES struct {
	Clock NMD_N02_CLOCK_AND_STATS_WITH_NOTES_CLOCK `hl7:"GROUP;ATR=optional"`
	AppStats NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATS `hl7:"GROUP;ATR=optional"`
	AppStatus NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATUS `hl7:"GROUP;ATR=optional"`
}

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES_CLOCK - Group struct
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES_CLOCK struct {
	SystemClock NCK `hl7:"TAG=NCK"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATS - Group struct
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATS struct {
	Statistics NST `hl7:"TAG=NST"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATUS - Group struct
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATUS struct {
	StatusChange NSC `hl7:"TAG=NSC"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMD_N02 - Network management data message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/NMD_N02
type NMD_N02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ClockAndStatsWithNotes []NMD_N02_CLOCK_AND_STATS_WITH_NOTES `hl7:"GROUP"`
}

