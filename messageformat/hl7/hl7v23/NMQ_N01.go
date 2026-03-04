package hl7v23

// NMQ_N01_QRY_WITH_DETAIL - Group struct
type NMQ_N01_QRY_WITH_DETAIL struct {
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
}

// NMQ_N01_CLOCK_AND_STATISTICS - Group struct
type NMQ_N01_CLOCK_AND_STATISTICS struct {
	SystemClock NCK `hl7:"TAG=NCK;ATR=optional"`
	Statistics NST `hl7:"TAG=NST;ATR=optional"`
	StatusChange NSC `hl7:"TAG=NSC;ATR=optional"`
}

// NMQ_N01 - Network management query message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/NMQ_N01
type NMQ_N01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QryWithDetail NMQ_N01_QRY_WITH_DETAIL `hl7:"GROUP;ATR=optional"`
	ClockAndStatistics []NMQ_N01_CLOCK_AND_STATISTICS `hl7:"GROUP"`
}

