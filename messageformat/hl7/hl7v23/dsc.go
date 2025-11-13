package hl7v23

type DSC struct {
	ContinuationPointer string `hl7:"POS=2" json:"continuationPointer,omitempty"`
}
