package functions

import (
	"testing"
	"time"

	"github.com/blutspende/go-parser/models"
	"github.com/blutspende/go-parser/parserconfig"
)

// Configuration structs for tests
var config *parserconfig.Configuration
var configHL7 *parserconfig.Configuration

// Reset configs to default values
func teardown() {
	config = parserconfig.NewDefaultConfiguration(parserconfig.ASTM)
	_ = parserconfig.InitConfig(config)
	configHL7 = parserconfig.NewDefaultConfiguration(parserconfig.HL7)
	_ = parserconfig.InitConfig(configHL7)
}

// Setup mock data for every test
func TestMain(m *testing.M) {
	// Set up configuration
	teardown()
	// Run all tests
	m.Run()
}

// Structure annotation helper factory
func createStructAnnotation(name string) models.StructAnnotation {
	return models.StructAnnotation{
		Tag: name,
	}
}

// Common test structures

// Annotation records
type Record struct {
	Field string `astm:"POS=3" hl7:"POS=3"`
}
type ComplexRecord struct {
	Field string `astm:"POS=3.2;ATR=length:4"`
}
type Group struct {
	Record Record `astm:"TAG=R" hl7:"TAG=REC"`
}
type Substructure struct {
	FirstComponent  string `astm:"POS=1" hl7:"POS=1"`
	SecondComponent string `astm:"POS=2" hl7:"POS=2"`
}

// Records
type ThreeFieldRecord struct {
	First  string `astm:"POS=3" hl7:"POS=3"`
	Second string `astm:"POS=4" hl7:"POS=4"`
	Third  string `astm:"POS=5" hl7:"POS=5"`
}
type SimpleRecord struct {
	First string `astm:"POS=3" hl7:"POS=3"`
}
type UnorderedRecord struct {
	First  string `astm:"POS=3"`
	Third  string `astm:"POS=5"`
	Second string `astm:"POS=4"`
}
type MultitypeRecord struct {
	String  string    `astm:"POS=3"`
	Int     int       `astm:"POS=4"`
	Float32 float32   `astm:"POS=5"`
	Float64 float64   `astm:"POS=6"`
	Time    time.Time `astm:"POS=7"`
}
type FloatLengthRecord struct {
	Default    float64 `astm:"POS=3"`
	Length0    float64 `astm:"POS=4;ATR=length:0"`
	Length4    float64 `astm:"POS=5;ATR=length:4"`
	LengthFull float64 `astm:"POS=6;ATR=length:-1"`
}
type MultitypePointerRecord struct {
	String  *string    `astm:"POS=3"`
	Int     *int       `astm:"POS=4"`
	Float32 *float32   `astm:"POS=5"`
	Float64 *float64   `astm:"POS=6"`
	Time    *time.Time `astm:"POS=7"`
}
type ComponentedRecord struct {
	First       string `astm:"POS=3"`
	SecondComp1 string `astm:"POS=4.1"`
	SecondComp2 string `astm:"POS=4.2"`
	ThirdComp1  string `astm:"POS=5.1"`
	ThirdComp2  string `astm:"POS=5.2"`
	ThirdComp3  string `astm:"POS=5.3"`
}
type ArrayRecord struct {
	First string   `astm:"POS=3"`
	Array []string `astm:"POS=4"`
}
type HeaderRecord struct {
	First string `astm:"POS=3" hl7:"POS=3"`
}
type HeaderDelimiterChange struct {
	First string   `astm:"POS=3"`
	Array []string `astm:"POS=4"`
	Comp1 string   `astm:"POS=5.1"`
	Comp2 string   `astm:"POS=5.2"`
}
type RequiredFieldRecord struct {
	First  string `astm:"POS=3"`
	Second string `astm:"POS=4;ATR=required"`
	Third  string `astm:"POS=5"`
}
type RequiredComponentRecord struct {
	First  string `astm:"POS=3.1"`
	Second string `astm:"POS=3.2;ATR=required"`
	Third  string `astm:"POS=3.3"`
}
type RecordType1 struct {
	First  string `astm:"POS=3"`
	Second int    `astm:"POS=4"`
}
type RecordType2 struct {
	First  int    `astm:"POS=3"`
	Second string `astm:"POS=4"`
}
type SubnameRecordType1 struct {
	Subname string `astm:"POS=3"`
	First   string `astm:"POS=4"`
	Second  int    `astm:"POS=5"`
}
type SubnameRecordType2 struct {
	Subname string `astm:"POS=3"`
	First   int    `astm:"POS=4"`
	Second  string `astm:"POS=5"`
}
type EnumString string
type EnumRecord struct {
	Enum EnumString `astm:"POS=3"`
}
type ReservedFieldRecord struct {
	TypeName  string `astm:"POS=1"`
	SeqNumber string `astm:"POS=2"`
}
type SparseFieldRecord struct {
	Field3 string `astm:"POS=3"`
	Field5 string `astm:"POS=5"`
}
type SubstructureField struct {
	FirstComponent  string `astm:"POS=1"`
	SecondComponent string `astm:"POS=2"`
	ThirdComponent  string `astm:"POS=3"`
}
type SubstructureRecord struct {
	First  string            `astm:"POS=3"`
	Second SubstructureField `astm:"POS=4"`
	Third  string            `astm:"POS=5"`
}
type SubstructureArrayRecord struct {
	First  string              `astm:"POS=3"`
	Second []SubstructureField `astm:"POS=4"`
	Third  string              `astm:"POS=5"`
}
type SparseSubstructureField struct {
	Component1 string `astm:"POS=1"`
	Component3 string `astm:"POS=3"`
	Component6 string `astm:"POS=6"`
}
type SparseSubstructureRecord struct {
	First  string                  `astm:"POS=3"`
	Second SparseSubstructureField `astm:"POS=4"`
}
type TimeRecord struct {
	Time time.Time `astm:"POS=3"`
}
type DateRecord struct {
	Date time.Time `astm:"POS=3;ATR=date"`
}
type WrongComponentOrderRecord struct {
	First string `astm:"POS=3"`
	Comp2 string `astm:"POS=4.2"`
	Comp1 string `astm:"POS=4.1"`
	Comp3 string `astm:"POS=4.3"`
}
type WrongComponentPlacementRecord struct {
	Field1 string `astm:"POS=3"`
	Comp1  string `astm:"POS=4.1"`
	Field2 string `astm:"POS=5"`
	Comp2  string `astm:"POS=4.2"`
}
type MultipleWrongComponentPlacementRecord struct {
	Field3 string `astm:"POS=3"`
	Comp41 string `astm:"POS=4.1"`
	Field5 string `astm:"POS=5"`
	Comp62 string `astm:"POS=6.2"`
	Comp42 string `astm:"POS=4.2"`
	Field7 string `astm:"POS=7"`
	Comp61 string `astm:"POS=6.1"`
	Field8 string `astm:"POS=8"`
}
type MissingAnnotationRecord struct {
	Field3  string `astm:"POS=3"`
	Missing string
	Field4  string `astm:"POS=4"`
}
type InvalidAttributeValueRecord struct {
	First float64 `astm:"POS=3;ATR=length:one"`
}

type SubSubField struct {
	First  string `astm:"POS=1" hl7:"POS=1"`
	Second string `astm:"POS=2" hl7:"POS=2"`
	Third  string `astm:"POS=3" hl7:"POS=3"`
}
type SubField struct {
	First  string      `astm:"POS=1" hl7:"POS=1"`
	Second SubSubField `astm:"POS=2" hl7:"POS=2"`
	Third  string      `astm:"POS=3" hl7:"POS=3"`
}
type SubSubRecord struct {
	First  string   `astm:"POS=3" hl7:"POS=3"`
	Second SubField `astm:"POS=4" hl7:"POS=4"`
	Third  string   `astm:"POS=5" hl7:"POS=5"`
}
type ReservedFieldRecordHL7 struct {
	One   string `hl7:"POS=1"`
	Two   string `hl7:"POS=2"`
	Three string `hl7:"POS=3"`
}
type NotReservedFieldRecordHL7 struct {
	Two   string `hl7:"POS=2"`
	Three string `hl7:"POS=3"`
}
type SequenceHl7 struct {
	Sequence int    `hl7:"POS=3;ATR=sequence"`
	Data     string `hl7:"POS=4"`
}

// Structures
type SingleRecordStruct struct {
	FirstRecord ThreeFieldRecord `astm:"TAG=R"`
}
type RecordArrayStruct struct {
	RecordArray []ThreeFieldRecord `astm:"TAG=R"`
}
type CompositeRecordStruct struct {
	Record1 RecordType1 `astm:"TAG=F"`
	Record2 RecordType2 `astm:"TAG=S"`
}
type CompositeMessage struct {
	CompositeRecordStruct CompositeRecordStruct `astm:"GROUP"`
}
type CompositeArrayMessage struct {
	CompositeRecordArray []CompositeRecordStruct `astm:"GROUP"`
}
type CompositeArrayAndSingleRecordMessage struct {
	CompositeRecordArray []CompositeRecordStruct `astm:"GROUP"`
	Ending               SimpleRecord            `astm:"TAG=E"`
}
type OptionalMessage struct {
	First    SimpleRecord `astm:"TAG=F"`
	Optional SimpleRecord `astm:"TAG=S;ATR=optional"`
	Third    SimpleRecord `astm:"TAG=T"`
}
type OptionalArrayMessage struct {
	First    SimpleRecord   `astm:"TAG=F"`
	Optional []SimpleRecord `astm:"TAG=A;ATR=optional"`
	Last     SimpleRecord   `astm:"TAG=L"`
}
type OptionalArrayAtTheEndMessage struct {
	First    SimpleRecord   `astm:"TAG=F"`
	Optional []SimpleRecord `astm:"TAG=A;ATR=optional"`
}
type OptionalAtTheEndMessage struct {
	First    SimpleRecord `astm:"TAG=F"`
	Optional SimpleRecord `astm:"TAG=O;ATR=optional"`
}
type SubnameMessage struct {
	Record1 SubnameRecordType1 `astm:"TAG=R;ATR=subname:FIRST"`
	Record2 SubnameRecordType2 `astm:"TAG=R;ATR=subname:SECOND"`
}
type SubnameArrayMessage struct {
	Array   []SubnameRecordType1 `astm:"TAG=R;ATR=subname:FIRST"`
	Record2 SubnameRecordType2   `astm:"TAG=R;ATR=subname:SECOND"`
}
type SubnameMultiArrayMessage struct {
	Array1 []SubnameRecordType1 `astm:"TAG=R;ATR=subname:FIRST"`
	Array2 []SubnameRecordType2 `astm:"TAG=R;ATR=subname:SECOND"`
}
type SubnameOptionalMessage struct {
	Record1 SubnameRecordType1 `astm:"TAG=R;ATR=subname:FIRST,optional"`
	Record2 SubnameRecordType2 `astm:"TAG=R;ATR=subname:SECOND"`
}
