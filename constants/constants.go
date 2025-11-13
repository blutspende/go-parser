package constants

// Maximum recurtion depth for message structures
const MaxDepth int = 42

// Annotation element types
const AnnotationElementGroup string = "GROUP"   // Composite message fragment structure
const AnnotationElementTag string = "TAG"       // Record identifier tag "TAG=R"
const AnnotationElementPosition string = "POS"  // Field position in record "POS=3"
const AnnotationElementAttribute string = "ATR" // Group, Record or Field attributes "ATR=required,length:5"

// Attributes for annotations
const AttributeOptional string = "optional" // Record or Group attribute for INPUT (by default all required)
const AttributeSubname string = "subname"   // Record attribute for specific subname "subname:MATRIX"
const AttributeRequired string = "required" // Field attribute for INPUT (by default all optional)
const AttributeDate string = "date"         // Field time.Time to OUTPUT as date only (YYYYMMDD)
const AttributeLength string = "length"     // Field float OUTPUT decimal length "length:2"
const AttributeSequence string = "sequence" // Field marked as sequence number in hl7
