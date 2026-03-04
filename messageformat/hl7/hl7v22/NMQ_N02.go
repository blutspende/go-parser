package hl7v22

// NMQ_N02_QRY_WITH_DETAIL - Group struct
type NMQ_N02_QRY_WITH_DETAIL struct {
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Clock_and_statistics []NMQ_N02_QRY_WITH_DETAIL_CLOCK_AND_STATISTICS `hl7:"GROUP"`
}

// NMQ_N02_QRY_WITH_DETAIL_CLOCK_AND_STATISTICS - Group struct
type NMQ_N02_QRY_WITH_DETAIL_CLOCK_AND_STATISTICS struct {
	SystemClock NCK `hl7:"TAG=NCK;ATR=optional"`
	Statistics NST `hl7:"TAG=NST;ATR=optional"`
	StatusChange NSC `hl7:"TAG=NSC;ATR=optional"`
}

// NMQ_N02 - Network Management Query
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/NMQ_N02
type NMQ_N02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	Qry_with_detail NMQ_N02_QRY_WITH_DETAIL `hl7:"GROUP;ATR=optional"`
}

