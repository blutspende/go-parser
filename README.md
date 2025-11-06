# go-parser
Library for handling ASTM and HL7 protocols in Go.

##### Install
`go get github.com/blutspende/go-parser`

# Features
  - Marshalling and unmarshalling of ASTM and HL7 messages
  - Baseline message structures for LIS02-A2, HL7 v2.3 and v2.4
  - Robust annotation system for defining custom message and record structures
  - Wide range of customizations of behaviour with a unified configuration parameter
  - Identifying message types without decoding
  - Automatic detection of line break type and custom protocol delimiters
  - Converting from-to raw bytes supporting many encodings and automatic timezone conversion

3 main functions and a utility is provided:
- `Marshal`: Converts a Go structure to an array of byte arrays
- `Unmarshal`: Converts a byte array to a Go structure
- `IdentifyMessage[ASTM/HL7]`: Identifies the type of message without decoding it
- `NewDefaultConfiguration`: Returns a copy of the default configuration for the given protocol
``` go
func Marshal(sourceStruct interface{}, config *pconfig.Configuration) (result [][]byte, err error) 
func Unmarshal(messageData []byte, targetStruct interface{}, config *pconfig.Configuration) (err error)
func IdentifyMessageASTM(messageData []byte, config *pconfig.Configuration) (messageType messagetype.MessageType, err error)
func IdentifyMessageHL7(messageData []byte, config *pconfig.Configuration) (messageType string, protocolVersion string, err error)
func NewDefaultConfiguration(protocol pconfig.Protocol) *pconfig.Configuration
```

# General information
In this section we provide an overview and some specific details about the whole message processing. This information is partially covered in other sections, where technical specifications are detailed, but here the goal is a general understanding of the concepts, logic and things to pay attention to.

## Message structures
The messages (source and target alike), which are often called "annotated structures" throughout this document, are regular go struct type definitions with specific [annotations](#annotated-structures) for each field. These are the bases for the whole parsing, provided directly as a parameter to Marshal and as a pointer parameter to write into to Unmarshal.

### Baseline
These structures can be built for each project using this library, but a baseline is provided following official protocol version documentations in this library. These structures can be found in the `messageformat` folder, under folders of `protocol/version` nesting. They can be directly used as is, mixed with user defined, copied and modified, or just as a guideline of the syntax for creating user defined structures. Any of these approaches work.

### Nesting
These structures however have a hierarchical nesting logic to mirror the complexities of an actual ASTM or HL7 message.
- Multi-message (optional): represents a repeated message of the same type
  - Content: single array of a message type
- Message: a structure of records and/or groups and/or arrays of them. Usually has a header record as the first item and potentially some terminator record at the end.
  - Content: Record(s) and/or Group(s) and/or arrays of them in any combination
- Group (optional): syntactically very similar to a message, without header. A represents a subsection of the message that can also be repeated if used as an array.
  - Content: Record(s) and/or Group(s) and/or arrays of them in any combination
- Record: a record always represents a single line in the message text, with a tag as the first field and all the actual data after. As the message consists of lines, the record is the first non-abstract data layer.
  - Content: Fields (matched to field separator separated actual primitive datatypes or substructures and/or arrays of them)
- Substructure: a Field can also have components which can either be defined as fields in the record with special annotation, or by using a separate struct type definition (a substructure)
  - Content: component fields (any primitive data type or sub-substructure)
- Sub-substructure (optional): in HL7 a substructure can have 1 more layer of substructure nested in it
  - Content: component fields (any primitive data type)

### Recursion
The structure nesting is processed recursively both in parsing and building, thus many layers can be handled just the same.
This is true for the message element level nesting and the record substructure nesting as well. However, there are limitations to both of them. At the message level, theoretically the nesting could be any finite number deep, but the limitation here is **42**. Below that an error will be produced.
The substructure nesting however does not allow many layers, because of protocol limitation (each layer has to have its own separator character), so here the protocol defines a hard limit. ASTM allow **1** and HL7 **2** layers of substructure nesting. Below that also an error is produced.

## Supplementary explanations
This section discusses a few specific details that help in understanding the behaviour and proper usage of the library.

### Date and time
Date and time fields are represented as `time.Time` types in Go structures. During parsing and building, the timezone conversion is handled automatically based on the configuration parameter [TimeZone](#timezone). All the conversions follow the same logic: the string representation is in the local time (stored in the configuration) and all the `time.Time` values stored in the Go structures are in UTC. So the Unmarshal will *result* in proper UTC times in the target structure and Marshal *expects* the times to be in UTC in the source structure.

An exception from this time zone conversion is the short date format (YYYYMMDD), which only represents a date without time. This is also stored in a `time.Time`, but only the date part is non-zero. The reason why it is exempt from the time zone conversion, is because this could potentially lead to date values shifting by one day: `20251106 -> 20251105230000`. Therefore, short dates should be kept in "local time" (configured timezone). If having UTC is more important, this behaviour can be disabled by setting the configuration parameter [KeepShortDateTimeZone](#keepshortdatetimezone) to false.

Being a short or long date format is determined differently in Unmarshal and Marshal. In Unmarshal it is only based on the input string length (8 or 14 characters), while in Marshal it is based on the presence of the `date` [attribute](#date) in the field's annotation, having this attribute meaning the short format and the default being the long.

### Pointers and empties
Pointer types can be used in fields of Record or Substructure level annotated structures in order to be able to signify missing fields for numeric types, which would otherwise have a default value of zero, which is indistinguishable from an actual value of zero.
This is especially important for marshaling with short notation, where the omission of empty fields at the end of a line is desired, and having a zero value somewhere long after the last actual value would result in a long field (and component) separated string result with no values apart from the zero.
Pointers can be used for any other primitive type as well, but there is not really a proper reason do to so.

Strings are considered missing if they are an empty string `""` even without being pointer type.
Date times also have a `time.IsZero()` option to indicate missing value without using pointers.

### Escape characters
In both ASTM and HL7 protocols, certain characters have special meanings as delimiters or control characters. To include these characters as literal values in the message content, escape sequences have to be used. In both protocols, the escape character and the delimiters are part of the header record, and can be message specific, however in all the examples the defaults will be used for demonstration.
This escaping can be used both in input messages (unmarshal) and output messages (marshal), however in marshal it has to be explicitly enabled in the configuration.

#### ASTM
In ASTM the escape is very simple: a single escape character is added before the special character to be escaped. The default escape character is `\`.

| Escape Sequence | Description                      | Represents Character |
|-----------------|----------------------------------|----------------------|
| \\\|            | Field separator character        | \|                   |
| \\\\            | Repetition separator character   | \\                   |
| \\^             | Component separator character    | ^                    |
| \\&             | Escape character                 | &                    |

#### HL7
In HL7 the escape sequences are more complex, and follow a specific pattern. The default escape character is also `\`, and the escape sequences are enclosed within two escape characters mostly with single capital letter codes. Here the escape sequence does not contain the represented escaped character at all.

| Escape Sequence | Description                      | Represents Character |
|-----------------|----------------------------------|----------------------|
| \\.br\\         | Carriage return                  | \\r                  |
| \\F\\           | Field separator character        | \|                   |
| \\R\\           | Repetition separator character   | ~                    |
| \\S\\           | Component separator character    | ^                    |
| \\T\\           | Subcomponent separator character | &                    |
| \\E\\           | Escape character                 | \\                   |
| \\Xhh\\         | Hexadecimal character value      | any                  |

# Library functions

## Default configuration
The `NewDefaultConfiguration` function returns a pointer to a new copy of the default configuration (including default delimiters) for the protocol selected by the only parameter to the function. This can then be used to modify the configuration for specific use cases while leaving the rest as default. This is the safe way to get a configuration instance, and it is not recommended to create a configuration instance manually. Defaults aim to keep behaviour backwards compatible even with updated library version when new functionalities are introduced.
``` go
config, err := pconfig.NewDefaultConfiguration(pconfig.PROTOCOL)
if err != nil {
    log.Fatal(err)
}
config.SomeValue = "non default value"
parser.SomeFunction(parameter, config)
```

## Identifying a message
Identifies the type of message without fully decoding it. Message identification means different things for ASTM and HL7, therefore there are two separate functions for each protocol.

### IdentifyMessageASTM
Return values are options from `github.com/blutspende/bloodlab-common/messagetype` enum definitions, including the `Unidentified` which will be returned if the message doesn't fit any known type. Note that this will not produce an error.
Uses the
Usage example:
``` go
messageType, err := parser.IdentifyMessageASTM([]byte(message), config)
if err != nil {
    log.Fatal(err)
}
switch (messageType) {
    case messagetype.Result:
      ...
    default:
      ...
}
```

### IdentifyMessageHL7
Returns `messageType` and `protocolVersion` as raw strings, the options are not limited to enum values.
The values are extracted from the MSH segment of the HL7 message. If the message type is missing, an error is returned, however protocol version is allowed be empty without an error.
Usage example:
``` go
messageType, protocolVersion, err := parser.IdentifyMessageHL7([]byte(message), config)
if err != nil {
    log.Fatal(err)
}
if protocolVersion != "2.4" {
    log.Fatalf("Unsupported protocol version: %s", protocolVersion)
}
if (messageType) {
    case "OUL^R21":
      ...
    default:
      ...
}
```

## Parsing a message: Unmarshal
The function Unmarshal takes a raw encoded byte array input of a message, and parses it into an [annotated structure](#annotated-structures). It can parse both ASTM and HL7 type messages, with the correct configuration set. The exact behaviour is determined by the annotations in the target structure and the [configuration](#configuration-structure) parameter. It can also parse multi-message inputs.
``` go
config, _ := pconfig.NewDefaultConfiguration(pconfig.ASTM)
err := parser.Unmarshal(inputBytes, &targetStruct, config)
if err != nil {
    log.Fatal(err)
}
fmt.Println(targetStruct.Record.Field)
```

## Building a message: Marshal
The function Marshal converts an [annotated structure](#annotated-structures) to a raw encoded array of byte arrays. This process is logically the opposite of the Unmarshal with some differences in details. It can also build for ASTM and HL7 based on the configuration parameter. Each element in the result array represents a line of the message, and thus has no line breaks at the end. The exact way the output is generated depends on the input structure's annotations and the [configuration](#configuration-structure).
``` go
config, _ := pconfig.NewDefaultConfiguration(pconfig.HL7)
lines, err := parser.Marshal(sourceStruct, config)
if err != nil {
    log.Fatal(err)
}
for _, line := range lines {
    fmt.Println(string(line))
}
```

# Configuration structure
For all three functions a configuration structure must be provided to determine behaviour.
``` go
type Configuration struct {
    Protocol                   pconfig.Protocol
    Encoding                   encoding.Encoding
    TimeZone                   timezone.TimeZone
    Notation                   string
    LineSeparator              string
    AutoDetectLineSeparator    bool
    EnforceSequenceNumberCheck bool
    DefaultDecimalPrecision    int
    RoundLastDecimal           bool
    KeepShortDateTimeZone      bool
    EscapeOutputStrings        bool
    EnforceMessageCompleteness bool
    DropEmptyOutputRecords     bool
    Delimiters                 Delimiters
    TimeLocation               *time.Location
}
```
Defaults for each protocol (including default delimiters) can be observed in `pconfig/models.go`.
``` go
DefaultConfigurationASTM
DefaultConfigurationHL7
```

## Protocol
Protocol selects from a predefined set of supported protocols (ASTM, HL7 as of now) and determines which protocol specific rules are applied during marshal and unmarshal. A mismatch of config protocol and the actual message will lead to errors in most cases.
Options are as defined in `pconfig/models.go`:
``` go
pconfig.ASTM
pconfig.HL7
```

## Encoding
Character encoding for reading and writing bytes. Options are all enum constants from `github.com/blutspende/bloodlab-common/encoding`.

## TimeZone
The timezone is used for date/time conversion. Options are all enum constants from `github.com/blutspende/bloodlab-common/timezone`.

## Notation
The notation is only used marshal. The notation is set to one of the following:
``` go
notation.Standard
notation.Short
```
Standard notation will produce as many fields as there are in the source structure, while short notation will omit empty fields at the end of a line. This is only relevant for marshal.

## LineSeparator
Line separator can be auto-detected, or set manually. If `AutoDetectLineSeparator` is set to true, this can be ignored. A few constants are provided for convenience, but any string is valid. This is only relevant for unmarshal.
``` go
lineseparator.LF
lineseparator.CR
lineseparator.LFCR
lineseparator.CRLF
```

## AutoDetectLineSeparator
If set to true, the line separator is detected automatically. If set to false, the line separator set in `LineSeparator` is used. This is only relevant for unmarshal.

## EnforceSequenceNumberCheck
In unmarshal, the sequence number (second field of every line in ASTM, and as defined with attributes in HL7) is checked for validity. If set to true, an error is returned if the sequence number is incorrect. If set to false, it is ignored. This is only relevant for unmarshal.

## DefaultDecimalPrecision
The default decimal precision is used for floating point numbers. If a field is not attributed with `length:N`, the default decimal precision is used. `-1` can be set to allow numbers to take any length required. This is only relevant for marshal.

## RoundLastDecimal
If it is set to true, floating point numbers are rounded up or down (based on common rounding rules) at the last decimal place determined by either `DefaultDecimalPrecision` or `length:N` annotation. If it is set to false the excess decimals are truncated. This is only relevant for marshal.

## KeepShortDateTimeZone
As short dates are only year, month, day, representing the date as midnight of that day in the `time.Time`, time zone conversions can lead to the change of day (e.g. 23:00 the day before). Because logically the short date represents a whole day, it can be more important to preserve the actual date as is then to have the time in UTC. However `time.Time` (and most database representations) must have a time zone, so the only solution is to keep the original time zone unconverted.
If this flag is set to true, the timezone is kept in local time for the short date format. If set to false, the time is converted to UTC just like long dates. This applies both for marshal and unmarshal (although what determines a date being short format differs), so with the same configuration the string format of the date will be intact.

## EscapeOutputStrings
If set to true, the output strings are escaped according to protocol rules and delimiters. If set to false, the output strings are not escaped, and will be output directly even if they contain delimiters. Default is false. This is only relevant for marshal.

## EnforceMessageCompleteness
If set to false, it is allowed - no error - for an input message to end prematurely with non-optional records or groups still unparsed in the target structure. This is only relevant for unmarshal.

## DropEmptyOutputRecords
If set to true, records that have all data fields empty (the name tag of the line is set by the annotation and therefore not considered) are dropped from the output. If set to false, all records are outputted, and empty (tag only) records will have a single field separator character added at the end. This is only relevant for marshal, and only in HL7 case (as ASTM always has the sequence number generated, thus tag only lines are not possible).

## Delimiters
Used for building the protocol's record structure. The defaults are part of the default config, set when calling `NewDefaultConfiguration`, but can be manually overridden if needed. Each field should contain exactly one character. This is mainly relevant for marshal. Unmarshal automatically detects the delimiters in the header record and stores it in this structure for the processing. 
``` go
type Delimiters struct {
    Field        string
    Repeat       string
    Component    string
    SubComponent string
    Escape       string
}
```

## TimeLocation
For internal use only. Should be ignored.

# Annotated structures
In order to read or write a message, an annotated structure is required. In a structure like this every field has to have an annotation with the protocol name and some information after that, like in this example:
``` go
type Structure struct {
    Field Type `astm:"information"`
}
```
The annotation tag name has to match the protocol we are using (and is also set in the configuration), and has to be all lower case. So the only valid options are:
``` go
astm
hl7
```

In all the [nested structures](#message-structures) syntactically there are two distinct types of structures:
- Message or Group: This is the structure that represents multiple lines, or the entire message. Its fields consist of records or other groups.
- Record or Substructure: This is the structure that represents a single line in the message, or parts of it. Its fields are actual primitive data types or substructures.

## Key-value system
Inside the quotations after the protocol is the actual information, which is based on a key-value system. They key-value pairs (if there are more than one) are listed with a separator character. Special attributes can optionally be added, which work the same way as a nested block with different separator characters.

### Syntax
At the top level the keys are always all capital letters, the list separator is a semicolon **;**, and the part separator is an equal sign **=**.
``` go
KEY1=value1;KEY2=value2
```
At the attribute level the keys are always lowercase letters, the list separator is a comma **,**, and the part separator is a colon **:**.
``` go
key1:value1,key2=value2
```

### Accepted keys
There are a limited number of keys accepted in each role. The raw definition can be seen in `constants/contstants.go`, but in the next chapter a detailed list and explanation is provided.

## Top level keys
| Key   | Value      | Where                  | On what | Explanation                                                         |
|-------|------------|------------------------|---------|---------------------------------------------------------------------|
| TAG   | name       | Messages, Groups       | Record  | the "name" of the record line (first field value) identifying it    |
| GROUP |            | Messages, Groups       | Group   | indicates the field is not a record, but a subgroup to recurse into |
| POS   | position   | Records, Substructures | Field   | position of the raw data field or substructure numbered from 1      |
| ATR   | attributes | any                    | any     | a list of attributes as key-value pairs                             |

### TAG and GROUP
On a Record field (inside a Message or Group) there has to be either a TAG or a GROUP annotation to determine its type.
The GROUP is just a flag, but the TAG has to have a value that matches the first field in the line which it tries to match to. It is a single capital letter in ASTM and a 3 character capital letter and number combination in HL7.
Examples:
``` go
PatientOrders []PatientOrder `astm:"GROUP"`

Manufacturer Manufacturer `astm:"TAG=M"`

PatientVisit PV1 `hl7:"TAG=PV1"
```

### POS
The POS is mandatory on each field in a Record or Substructure. This is what positions the data elements in a record line. All the numbering start with 1. In Substructures, it is just a single integer, however in Records it is allowed to have multiple fields pointing to the same Field, being components in it, basically a first level substructure not defined as separate type. In this case the field and component positions are separated by a dot.
Examples:
``` go
NoticeOfAdmissionCodestring    `hl7:"POS=24"`
NoticeOfAdmissionDatetime.Time `hl7:"POS=25"`

DataMeasurementValue     string `astm:"POS=4.1"`
InitialMeasurementValue  string `astm:"POS=4.2"`
MeasurementValueOfDevice string `astm:"POS=4.3"`
```
In ASTM the first two fields (position 1 and 2) are reserved for the record name (TAG) and the sequence number, so the actual data fields can only start from position 3. In HL7 only the first field (position 1) is reserved for the record name (TAG).

### ATR
The ATR attributes have the value of a whole nested key-value list, as detailed below.

## Attribute level keys
| Key       | Value  | Where                   | On what   | Explanation                                                    |
|-----------|--------|-------------------------|-----------|----------------------------------------------------------------|
| optional  |        | Messages, Groups        | any       | input record or section can be skipped in not found in message |
| subname   | name   | Messages, Groups        | Record    | additional identifier of the Record in the third field         |
| required  |        | Records, Substructures  | Field     | input is mandatory                                             |
| date      |        | Records, Substructures  | time.Time | output as date only (YYYYMMDD)                                 |
| length    | number | Records, Substructures  | float     | output decimal length                                          |
| sequence  |        | Records, Substructures  | int       | marked as sequence number in hl7                               |

### optional
By default, at the message level all elements are mandatory, so the 'optional' has to be added as an attribute. In case of a Group an optional attribute will recursively trickle down to all its elements becoming effectively optional even if it is not explicitly declared for them.
``` go
Patient Patient `hl7:"GROUP;ATR=optional"`
```

### subname
Subname can be used in special cases for Records that have the same TAG but still differ in structure. Some instruments send such messages adding an extra indicator in the third field, which will be matched with the subname attribute value, and then behave exactly like it would be a TAG match or mismatch for the message structure processing.
``` go
Histograms []Stream `astm:"TAG=M;ATR=subname:HISTOGRAM"`
Matrices   []Stream `astm:"TAG=M;ATR=subname:MATRIX"`
```

### required
By default, at the record field level all fields are optional, so the 'required' attribute has to be added to make it mandatory, and thus produce an error if the field is empty.

### date
The 'date' attribute can be used on time.Time fields to signify that only the date part is relevant, and the time part should be ignored. In Marshal, it will be outputted in short date format (YYYYMMDD). It will also forego time zone conversion (if KeepShortDateTimeZone is set to true in the configuration), to avoid the actual date shifting with one day because it became "11pm previous day" or similar. Unmarshal will process depending on the input format (8 or 14 characters), so the attribute does not affect it.
``` go
PlanEffectiveDate time.Time `hl7:"POS=13;ATR=date"
```

### length
The 'length' attribute can be used on float fields (float32, float64) to signify the number of decimals to output. The value has to be an integer >= -1. If -1 is set, the number will be outputted with maximum number of decimals. If >=0 is set, the number will be outputted with exactly that many decimals, either truncating or rounding the excess decimals based on configuration.
``` go
MeasuredValue float32 `astm:"POS=5;ATR=length:2"`
```

### sequence
The 'sequence' attribute can be used on int fields to signify that this field is the sequence number of the record in HL7 messages. This will make the field automatically incremented in marshal, and checked for validity in unmarshal (if enabled in configuration). ASTM has a fixed position (field 2) for the sequence number, so this attribute is not relevant there.
``` go
SequenceNumber int `hl7:"POS=2;ATR=sequence"`
```

## Record examples
Here are a few examples of annotated records and their corresponding message lines.

### Array
``` go
type Record struct {
    Field1 []string `astm:"POS=3"`
}
```
```
R|1|value1\value2\value3
```

### Components
``` go
type Record struct {
    Field1Component1 string `astm:"POS=3.1"`
    Field1Component2 string `astm:"POS=3.2"`
    Field1Component3 string `astm:"POS=3.3"`
}
```
```
R|1|component1^component2^component3
```

### Substructure
``` go
type Substructure struct {
    Component1 string `astm:"POS=1"`
    Component2 string `astm:"POS=2"`
}
type Record struct {
    Field1 Substructure `astm:"POS=3"`
}
```
```
R|1|component1^component2
```

### Substructure array
``` go
type Substructure struct {
    Component1 string `astm:"POS=1"`
    Component2 string `astm:"POS=2"`
}
type Record struct {
    Field1 []Substructure `astm:"POS=3"`
}
```
```
R|1|comp1^comp2\comp1^comp2\comp1^comp2
```

## Message examples
Here are a few messages, groups and arrays of them with annotations and their possible corresponding message lines.

### Simple message
``` go
type Lis02a2Message {
    MessageHeader lis02a2.Header     `astm:"H"`
    Record        RecordType         `astm:"R"`
    Terminator    lis02a2.Terminator `astm:"L"`
}
```
```
H|\^&||||
R|1|value1|value2
L|1|N
```

### Record array
``` go
type HL7v23Message {
	MSH    hl7v23.MSH   `hl7:"TAG=MSH"`
	Record []RecordType `hl7:"REC"`
}
```
```
MSH|^~\&|
REC|value1|value2
REC|value1|value2
REC|value1|value2
```

### Group
``` go
type GroupType {
    Record1 FirstRecordType  `hl7:"RT1"`
    Record2 SecontRecordType `hl7:"RT2"`
}
type HL7v23Message {
	MSH   hl7v23.MSH `hl7:"TAG=MSH"`
    Group GroupType  `hl7:"GROUP"`
}
```
```
MSH|^~\&|
RT1|1|value1
RT2|value2|20251106|3.14159
```

### Group array
``` go
type Lis02a2Message {
    MessageHeader lis02a2.Header     `astm:"TAG=H"`
    GroupArray    []GroupType        `astm:"GROUP"`
    Terminator    lis02a2.Terminator `astm:"TAG=L"`
}
```
```
H|\^&||||
F|1|value1|value2
S|1|value1|value2
F|2|value1|value2
S|1|value1|value2
F|3|value1|value2
S|1|value1|value2
L|1|N
```
Note that the sequence number is incremented for each instance of the nested structure, however only the first record of the nested structure takes the sequence number, and the rest is 1 (unless the nested structure has its own array inside).