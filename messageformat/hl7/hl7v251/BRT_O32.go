package hl7v251

// BRT_O32_RESPONSE - Group struct
type BRT_O32_RESPONSE struct {
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	Order []BRT_O32_RESPONSE_ORDER `hl7:"GROUP;ATR=optional"`
}

// BRT_O32_RESPONSE_ORDER - Group struct
type BRT_O32_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []BRT_O32_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	BloodProductOrder BPO `hl7:"TAG=BPO;ATR=optional"`
	BloodProductTransfusionDisposition []BTX `hl7:"TAG=BTX;ATR=optional"`
}

// BRT_O32_RESPONSE_ORDER_TIMING - Group struct
type BRT_O32_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// BRT_O32 - Blood Product Transfusion/Disposition Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/BRT_O32
type BRT_O32 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response BRT_O32_RESPONSE `hl7:"GROUP;ATR=optional"`
}

