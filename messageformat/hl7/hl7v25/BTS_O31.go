package hl7v25

// BTS_O31_PATIENT - Group struct
type BTS_O31_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit BTS_O31_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// BTS_O31_PATIENT_PATIENT_VISIT - Group struct
type BTS_O31_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// BTS_O31_ORDER - Group struct
type BTS_O31_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []BTS_O31_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	BloodProductOrder BPO `hl7:"TAG=BPO"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ProductStatus []BTS_O31_ORDER_PRODUCT_STATUS `hl7:"GROUP;ATR=optional"`
}

// BTS_O31_ORDER_TIMING - Group struct
type BTS_O31_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// BTS_O31_ORDER_PRODUCT_STATUS - Group struct
type BTS_O31_ORDER_PRODUCT_STATUS struct {
	BloodProductTransfusionDisposition BTX `hl7:"TAG=BTX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// BTS_O31 - Blood Product Transfusion/Disposition
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/BTS_O31
type BTS_O31 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient BTS_O31_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []BTS_O31_ORDER `hl7:"GROUP"`
}

