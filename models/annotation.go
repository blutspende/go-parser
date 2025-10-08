package models

// Annotation types for fields and structures

type FieldAnnotation struct {
	Raw            string
	FieldPos       int
	IsArray        bool
	IsComponent    bool
	ComponentPos   int
	IsSubstructure bool
	Attributes     map[string]string
}

type StructAnnotation struct {
	Raw         string
	StructName  string
	IsComposite bool
	IsArray     bool
	Attributes  map[string]string
}
