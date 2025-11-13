package hl7v24

import "time"

// HL7 v2.4 - OBX - Observation/Result
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OBX
type OBX struct {
	SetID                        int       `hl7:"POS=2;ATR=sequence" json:"SetID,omitempty"`
	ValueType                    string    `hl7:"POS=3" json:"ValueType,omitempty"`
	ObservationIdentifier        CE        `hl7:"POS=4" json:"ObservationIdentifier,omitempty"`
	ObservationSubID             string    `hl7:"POS=5" json:"ObservationSubID,omitempty"`
	ObservationValue             []string  `hl7:"POS=6" json:"ObservationValue,omitempty"`
	Units                        CE        `hl7:"POS=7" json:"Units,omitempty"`
	ReferenceRange               string    `hl7:"POS=8" json:"ReferenceRange,omitempty"`
	AbnormalFlags                []string  `hl7:"POS=9" json:"AbnormalFlags,omitempty"`
	Probability                  *float32  `hl7:"POS=10" json:"Probability,omitempty"`
	NatureOfAbnormalTest         []string  `hl7:"POS=11" json:"NatureOfAbnormalTest,omitempty"`
	ResultStatus                 string    `hl7:"POS=12" json:"ResultStatus,omitempty"`
	DateLastObservedNormalValues time.Time `hl7:"POS=13" json:"DateLastObservedNormalValues,omitempty"`
	UserDefinedAccessChecks      string    `hl7:"POS=14" json:"UserDefinedAccessChecks,omitempty"`
	DateTimeOfObservation        time.Time `hl7:"POS=15" json:"DateTimeOfObservation,omitempty"`
	ProducersID                  CE        `hl7:"POS=16" json:"ProducersID,omitempty"`
	ResponsibleObserver          XCN       `hl7:"POS=17" json:"ResponsibleObserver,omitempty"`
	ObservationMethod            []CE      `hl7:"POS=18" json:"ObservationMethod,omitempty"`
	EquipmentInstanceIdentifier  EI        `hl7:"POS=19" json:"EquipmentInstanceIdentifier,omitempty"`
	DateTimeOfTheAnalysis        time.Time `hl7:"POS=20" json:"DateTimeOfTheAnalysis,omitempty"`
}
