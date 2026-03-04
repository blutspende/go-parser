package hl7v25

// NMQ_N01_QRY_WITH_DETAIL - Group struct
type NMQ_N01_QRY_WITH_DETAIL struct {
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

// NMQ_N01_CLOCK_AND_STATISTICS - Group struct
type NMQ_N01_CLOCK_AND_STATISTICS struct {
	SystemClock NCK `hl7:"TAG=NCK;ATR=optional"`
	ApplicationControlLevelStatistics NST `hl7:"TAG=NST;ATR=optional"`
	ApplicationStatusChange NSC `hl7:"TAG=NSC;ATR=optional"`
}

// NMQ_N01 - Application management query message
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/NMQ_N01
type NMQ_N01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QryWithDetail NMQ_N01_QRY_WITH_DETAIL `hl7:"GROUP;ATR=optional"`
	ClockAndStatistics []NMQ_N01_CLOCK_AND_STATISTICS `hl7:"GROUP"`
}

