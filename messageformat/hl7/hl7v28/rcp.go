package hl7v28

import "time"

// RCP - Response Control Parameter
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/RCP
type RCP struct {
	QueryPriority string `hl7:"POS=2"`
	QuantityLimitedRequest CQ `hl7:"POS=3"`
	ResponseModality CNE `hl7:"POS=4"`
	ExecutionAndDeliveryTime time.Time `hl7:"POS=5"`
	ModifyIndicator string `hl7:"POS=6"`
	SortByField []SRT `hl7:"POS=7"`
	SegmentGroupInclusion []string `hl7:"POS=8"`
}

