package hl7v24

import "time"

// HL7 v2.4 - OBX - Observation/Result
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OBX
type OBX struct {
	SetID                        int       `hl7:"1,sequence" json:"SetID,omitempty"`
	ValueType                    string    `hl7:"2" json:"ValueType,omitempty"`
	ObservationIdentifier        CE        `hl7:"3" json:"ObservationIdentifier,omitempty"`
	ObservationSubID             string    `hl7:"4" json:"ObservationSubID,omitempty"`
	ObservationValue             []string  `hl7:"5" json:"ObservationValue,omitempty"`
	Units                        CE        `hl7:"6" json:"Units,omitempty"`
	ReferenceRange               string    `hl7:"7" json:"ReferenceRange,omitempty"`
	AbnormalFlags                []string  `hl7:"8" json:"AbnormalFlags,omitempty"`
	Probability                  float32   `hl7:"9" json:"Probability,omitempty"`
	NatureOfAbnormalTest         []string  `hl7:"10" json:"NatureOfAbnormalTest,omitempty"`
	ResultStatus                 string    `hl7:"11" json:"ResultStatus,omitempty"`
	DateLastObservedNormalValues time.Time `hl7:"12,longdate" json:"DateLastObservedNormalValues,omitempty"`
	UserDefinedAccessChecks      string    `hl7:"13" json:"UserDefinedAccessChecks,omitempty"`
	DateTimeOfObservation        time.Time `hl7:"14,longdate" json:"DateTimeOfObservation,omitempty"`
	ProducersID                  CE        `hl7:"15" json:"ProducersID,omitempty"`
	ResponsibleObserver          XCN       `hl7:"16" json:"ResponsibleObserver,omitempty"`
	ObservationMethod            []CE      `hl7:"17" json:"ObservationMethod,omitempty"`
	EquipmentInstanceIdentifier  EI        `hl7:"18" json:"EquipmentInstanceIdentifier,omitempty"`
	DateTimeOfTheAnalysis        time.Time `hl7:"19" json:"DateTimeOfTheAnalysis,omitempty"`
}
