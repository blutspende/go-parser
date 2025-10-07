package hl7v23

type DSC struct {
	ContinuationPointer string `hl7:"1" json:"continuationPointer,omitempty"`
}
