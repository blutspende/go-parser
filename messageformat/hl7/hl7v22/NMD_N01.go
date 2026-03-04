package hl7v22

// NMD_N01_CLOCK_AND_STATS_WITH_NOTES - Group struct
type NMD_N01_CLOCK_AND_STATS_WITH_NOTES struct {
	Clock NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK `hl7:"GROUP;ATR=optional"`
}

// NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK - Group struct
type NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK struct {
	SystemClock NCK `hl7:"TAG=NCK"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	App_stats NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK_APP_STATS `hl7:"GROUP;ATR=optional"`
}

// NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK_APP_STATS - Group struct
type NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK_APP_STATS struct {
	Statistics NST `hl7:"TAG=NST"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	App_status NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK_APP_STATS_APP_STATUS `hl7:"GROUP;ATR=optional"`
}

// NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK_APP_STATS_APP_STATUS - Group struct
type NMD_N01_CLOCK_AND_STATS_WITH_NOTES_CLOCK_APP_STATS_APP_STATUS struct {
	StatusChange NSC `hl7:"TAG=NSC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMD_N01 - Network Management Data
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/NMD_N01
type NMD_N01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	Clock_and_stats_with_notes []NMD_N01_CLOCK_AND_STATS_WITH_NOTES `hl7:"GROUP"`
}

