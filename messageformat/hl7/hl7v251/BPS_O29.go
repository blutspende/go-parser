package hl7v251

// BPS_O29_PATIENT - Group struct
type BPS_O29_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit BPS_O29_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// BPS_O29_PATIENT_PATIENT_VISIT - Group struct
type BPS_O29_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// BPS_O29_ORDER - Group struct
type BPS_O29_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []BPS_O29_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	BloodProductOrder BPO `hl7:"TAG=BPO"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Product []BPS_O29_ORDER_PRODUCT `hl7:"GROUP;ATR=optional"`
}

// BPS_O29_ORDER_TIMING - Group struct
type BPS_O29_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// BPS_O29_ORDER_PRODUCT - Group struct
type BPS_O29_ORDER_PRODUCT struct {
	BloodProductDispenseStatus BPX `hl7:"TAG=BPX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// BPS_O29 - Blood Product Dispense Status Message
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/BPS_O29
type BPS_O29 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient BPS_O29_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []BPS_O29_ORDER `hl7:"GROUP"`
}

