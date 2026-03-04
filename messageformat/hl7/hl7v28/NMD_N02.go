package hl7v28

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES - Group struct
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES struct {
	Clock NMD_N02_CLOCK_AND_STATS_WITH_NOTES_CLOCK `hl7:"GROUP;ATR=optional"`
	App_stats NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATS `hl7:"GROUP;ATR=optional"`
	App_status NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATUS `hl7:"GROUP;ATR=optional"`
}

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES_CLOCK - Group struct
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES_CLOCK struct {
	SystemClock NCK `hl7:"TAG=NCK"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATS - Group struct
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATS struct {
	ApplicationControlLevelStatistics NST `hl7:"TAG=NST"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATUS - Group struct
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES_APP_STATUS struct {
	ApplicationStatusChange NSC `hl7:"TAG=NSC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// NMD_N02 - Application management data message (unsolicited)
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/NMD_N02
type NMD_N02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Clock_and_stats_with_notes []NMD_N02_CLOCK_AND_STATS_WITH_NOTES `hl7:"GROUP"`
}

