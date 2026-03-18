# go-parser
Library for handling ASTM and HL7 protocols in Go.

##### Install
`go get github.com/blutspende/go-parser`

# Features
  - Marshaling and unmarshalling ASTM and HL7 messages
  - Predefined message structures for ASTM LIS02-A2 and HL7 v2.3, v2.4
  - Robust annotation system for defining custom message types
  - Wide range of behavior customizations with a unified configuration parameter
  - Identifying message types without decoding
  - Automatic detection of line break type and custom protocol delimiters
  - Converting from-to raw bytes supporting many encodings and automatic timezone conversion

Two main functions and two utilities are provided:
- `Unmarshal`: Converts a byte array to a Go structure
- `Marshal`: Converts a Go structure to an array of byte arrays
- `IdentifyMessage[ASTM/HL7]`: Identifies the type of message without decoding it
- `NewDefaultConfiguration`: Returns a copy of the default configuration for the given protocol
``` go
func Marshal(sourceStruct interface{}, config *pconfig.Configuration) (result [][]byte, err error) 
func Unmarshal(messageData []byte, targetStruct interface{}, config *pconfig.Configuration) (err error)
func IdentifyMessageASTM(messageData []byte, config *pconfig.Configuration) (messageType instrumentenum.MessageType, err error)
func IdentifyMessageHL7(messageData []byte, config *pconfig.Configuration) (messageType string, protocolVersion string, err error)
func NewDefaultConfiguration(protocol pconfig.Protocol) *pconfig.Configuration
```

# General information
This section provides an overview as well as some specific details. This information is partially covered in other more technical sections, here the goal is to give a general understanding and things to pay attention to.

## Message structures
The messages are mapped to or from Go structures, these are called source and target structures depending on the parsing direction (marshal or unmarshal). These are also referred to as "annotated structures" throughout this document, which are regular Go struct type definitions with specific [annotations](#annotated-structures) for each field. These are the bases for the whole parsing, provided directly as a parameter to Marshal and as a pointer parameter to Unmarshal.

### Baseline
The annotated structures can be specifically built for each user project, but a baseline is provided by this library following official protocol version documentations. They can be found in `messageformat/[protocol]/[version]` folders. They can be used directly as is, mixed with user-defined structures, copied and modified, or taken as a guideline for creating completely user-defined structures. Any of these approaches work.

### Nesting
These structures have a hierarchical nesting logic to mirror the complexities of an ASTM or HL7 message.
These are from top to bottom:
- Multi-message (optional): represents a repeated message of the same type
  - Content: single array of a message type
- Message: a structure of records and/or groups and/or arrays of them. Usually has a header record as the first item and potentially some terminator record at the end.
  - Content: Record(s) and/or Group(s) and/or arrays of them in any combination
- Group (optional): syntactically very similar to a message without header. Represents a subsection of the message that can also be repeated if used as an array.
  - Content: Record(s) and/or Group(s) and/or arrays of them in any combination
- Record: a record always represents a single line in the message text, with a tag as the first field. As the raw messages consist of lines, the record is the first non-abstract data layer.
  - Content: Fields (matched to field separator separated actual primitive datatypes or substructures and/or arrays of them)
- Substructure: a Field can also have components which can either be defined as fields in the record with special annotation, or by using a separate struct type definition (a substructure)
  - Content: component fields (any primitive data type or sub-substructure)
- Sub-substructure (optional): in HL7 a substructure can have one more layer of substructure nested in it
  - Content: component fields (any primitive data type)

### Recursion
The structure nesting is processed recursively both in parsing (unmarshal) and building (marshal), thus many layers can be handled with the same logic.

This is true for the message element level nesting and the record substructure nesting as well. However, there are limitations to both of them. At the message level, theoretically the nesting could be any finite number deep, but the limit here is **42**. Below that an error will be produced.

The substructure nesting does not allow many layers, because of protocol limitation (each layer has to have its own separator character), so here the protocol defines a hard limit. ASTM allow **1** and HL7 **2** layers of substructure nesting. Below that also an error is produced.

## Supplementary explanations
This section discusses a few specific details that help in understanding the behaviour and proper usage of the library.

### Date and time
Date and time fields are represented as `time.Time` types in Go structures. During parsing and building, the timezone conversion is handled automatically based on the configuration parameter [TimeZone](#timezone). All the conversions follow the same logic: the raw message string representation are considered to be in the "local time" zone (stored in the configuration) and the `time.Time` values stored in the Go structures are in UTC. So the Unmarshal will *result* in proper UTC times in the target structure and Marshal *expects* the times to be in UTC in the source structure.

An exception from this time zone conversion is the short date format (YYYYMMDD), which only represents a logical calendar date without time. This is also stored in a `time.Time`, but after the day part everything is zero. The reason why it can be exempt from the time zone conversion, is because this could potentially lead to date values shifting by one day: `20251106 -> 20251105230000` and thus ruining the date part. Therefore, short dates can be kept in "local time" (configured timezone) maintaining the year month and date exactly as in the string. If having UTC is more important, this behaviour can be disabled by setting the configuration parameter [KeepShortDateTimeZone](#keepshortdatetimezone) to false.

Date format is determined differently in Unmarshal and Marshal. In Unmarshal it is only based on the input string length (8 or 14 characters), while in Marshal it is based on the presence of the [date attribute](#date) in the field's annotation, having this attribute meaning the short format and the default being the long.

### Pointers and empties
Pointer types can be used in fields of Record and Substructure level annotated structures in order to be able to signify missing fields for numeric types, which would otherwise have a default value of zero, and be indistinguishable from an actual value of zero.
This is especially important for marshaling with short notation, where the omission of empty fields at the end of a line is desired, and having a zero value somewhere long after the last actual value would result in long meaningless strings with a few zeros.

Pointers can be used for any other primitive type as well, but there is not really a reason do to so.
Strings are considered missing if they are an empty string `""` even without being pointer type.
Date times also have a `time.IsZero()` option to indicate missing value without using pointers.

### Floating point precision
When building messages (marshal) floating point numbers can have a specific number of decimal places. This can either be set individually by the [length attribute](#length) in the field's annotation, or globally in the configuration parameter [DefaultDecimalPrecision](#defaultdecimalprecision). The individual field annotation takes precedence over the global configuration.
Trailing decimals will either be rounded or truncated based on the [RoundLastDecimal](#roundlastdecimal) configuration parameter.
The rules for the length value are as follows:

| Value   | Explanation                            |
|---------|----------------------------------------|
| N < -1  | invalid value, will produce an error   |
| N = -1  | output with maximum number of decimals |
| N = 0   | integer part only                      |
| N > 0   | output with exactly that many decimals |

### Escape characters
In both ASTM and HL7 protocols, certain characters have special meanings as delimiters or control characters. To include these characters as literal values in string fields, escape sequences have to be used. In both protocols, the escape character and the delimiters are part of the header record, and can be message specific, but in all the examples here the defaults will be used.
Escaping is automatically used for input messages (unmarshal) and has to be enabled in the [configuration](#escapeoutputstrings) for output messages (marshal).

#### ASTM
In ASTM the escape is very simple: a single escape character is added before the special character to be escaped. The default escape character is `&`.

| Escape Sequence | Description                      | Represents Character |
|-----------------|----------------------------------|----------------------|
| &\|             | Field separator character        | \|                   |
| &\\             | Repetition separator character   | \\                   |
| &^              | Component separator character    | ^                    |
| &&              | Escape character                 | &                    |

#### HL7
In HL7 the escape sequences are more complex, and follow a specific pattern. The default escape character is `\`, and the escape sequences are enclosed within two escape characters mostly with single capital letter codes. Here the escape sequence does not contain the represented characters at all.

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
Description and examples of the 4 main functions provided by the library.

## Default configuration
The `NewDefaultConfiguration` function returns a pointer to a new copy of the default configuration (including default delimiters) for the protocol selected by the only parameter to the function. This can then be used to modify the configuration for specific use cases while leaving the rest as default. This is the safe way to get a configuration instance, and it is not recommended to create a configuration instance manually. Defaults aim to keep behavior backwards compatible even after version update with new functionalities introduced.
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
Return values are options from `github.com/blutspende/bloodlab-common/messagetype` enum definitions, including the `Unidentified` which will be returned if the message doesn't fit any known type, but this will not produce an error. The type depends on the pattern of the records types in the message.

The identifier extracts a signature of the message, which is the first character (the name tag) of each record line, with `C` and `M` records being filtered out. It is then matched as a whole string against a predefined set of regex patterns. It is possible to inject [custom rules](#CustomASTMIdentifyRules) for identifying messages, using the `CustomASTMIdentifyRules` array in the configuration parameter.
``` go
messageType, err := parser.IdentifyMessageASTM(inputBytes, config)
if err != nil {
    log.Fatal(err)
}
switch (messageType) {
    case instrumentenum.Result:
      ...
    default:
      ...
}
```

### IdentifyMessageHL7
Returns `messageType` and `protocolVersion` as raw strings, the options are not limited to predefined values.
The values are extracted from the MSH segment of the HL7 message. If the message type is missing, an error is returned, but protocol version is allowed be missing without an error.
``` go
messageType, protocolVersion, err := parser.IdentifyMessageHL7(inputBytes, config)
if err != nil {
    log.Fatal(err)
}
if protocolVersion != "2.4" {
    log.Fatalf("unsupported protocol version: %s", protocolVersion)
}
if (messageType) {
    case "OUL^R21":
      ...
    default:
      ...
}
```

## Parsing a message: Unmarshal
The function Unmarshal takes the input message (a raw encoded byte array) and an [annotated structure](#annotated-structures) to map it into. It can parse both ASTM and HL7 type messages depending on the [protocol](#protocol) in the configuration. The exact behaviour is determined by the annotations in the target structure and the [configuration](#configuration-structure) parameter. It can also parse multi-message inputs.
``` go
config, _ := pconfig.NewDefaultConfiguration(pconfig.PROTOCOL)
err := parser.Unmarshal(inputBytes, &targetStruct, config)
if err != nil {
    log.Fatal(err)
}
fmt.Println(targetStruct.Record.Field)
```

## Building a message: Marshal
The function Marshal converts an [annotated structure](#annotated-structures) to a raw encoded array of byte arrays. This process is logically the opposite of the Unmarshal with some differences in details. It can also build for ASTM and HL7 depending on the [protocol](#protocol) in the configuration. Each element in the result array represents a line of the message, and thus has no line breaks at the end. The exact way the output is generated depends on the input structure's annotations and the [configuration](#configuration-structure).
``` go
config, _ := pconfig.NewDefaultConfiguration(pconfig.PROTOCOL)
lines, err := parser.Marshal(sourceStruct, config)
if err != nil {
    log.Fatal(err)
}
for _, line := range lines {
    fmt.Println(string(line))
}
```

# Configuration structure
For all three functions a configuration structure must be provided to determine behavior.
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
    CustomASTMIdentifyRules    []ASTMIdentifyRule
    TimeLocation               *time.Location
}
```
The defaults for each protocol (including default delimiters) can be found in `pconfig/models.go`.
``` go
DefaultConfigurationASTM
DefaultConfigurationHL7
```

## Protocol
Protocol is a string enum with a predefined set of supported protocols: ASTM and HL7 as of now. It determines which protocol specific rules are applied during marshal and unmarshal. A mismatch of config protocol and the actual message will lead to errors in most cases.
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
The default decimal precision is used for floating point numbers. Detailed explanation is provided in the [Floating point precision](#floating-point-precision) section. This is only relevant for marshal.

## RoundLastDecimal
If it is set to true, floating point numbers are rounded up or down (based on common rounding rules) at the last decimal place. If it is set to false the excess decimals are truncated. Detailed explanation is provided in the [Floating point precision](#floating-point-precision) section. This is only relevant for marshal.

## KeepShortDateTimeZone
If this flag is set to true, the timezone is kept in local time for the short date format. If set to false, the time is converted to UTC just like long dates. This applies both for marshal and unmarshal (although what determines a date being short format differs), so with the same configuration the string format of the date will be intact.
Why this is an issue is explained in the [Date and time](#date-and-time) section.

## EscapeOutputStrings
If set to true, the output strings are [escaped](#escape-characters) according to protocol rules and delimiters. If set to false, the output strings are not escaped, and will be output directly even if they contain delimiters. Default is false. This is only relevant for marshal.

## EnforceMessageCompleteness
If set to false, it is allowed - no error - for an input message to be incomplete and end prematurely with non-optional records or groups still unparsed in the target structure. This is only relevant for unmarshal.

## DropEmptyOutputRecords
If set to true, records (lines) that have all informative data fields [empty](#pointers-and-empties) are dropped from the output. Fields are considered informative that are manually set in the source structure, so tag fields (field 1) and sequence number fields (usually field 2) are automatic and do not count.
If set to false, all records are outputted no matter what. In this case a tag only record will have a single field separator character added at the end. This is only relevant for marshal, and only in HL7. ASTM does not allow empty records.

## Delimiters
Used for building the record's internal structure. The defaults are part of the default config, set when calling `NewDefaultConfiguration`, but can be manually overridden if needed. Each field should contain exactly one character. This is mainly relevant for marshal, as Unmarshal automatically detects the delimiters in the header record and stores it in this structure for the processing. 
``` go
type Delimiters struct {
    Field        string
    Repeat       string
    Component    string
    SubComponent string
    Escape       string
}
```

## CustomASTMIdentifyRules
Can be used to inject custom rules for identifying ASTM messages. It is only used in the `IdentifyMessageASTM` [function](#IdentifyMessageASTM). `CustomASTMIdentifyRules` is an array of `ASTMIdentifyRule` structs, where each struct represents a match rule. This field is `nil` by default and can be left empty if there are no custom rules required.
``` go
type ASTMIdentifyRule struct {
    Name  string
    Regex string
    Type  instrumentenum.MessageType
}
```
- `Name` is used for logging and better readability.
- `Regex` is a regular expression matched against the signature of the message as a whole string. The signature consists of the first character of each record line, with `C` and `M` records filtered out. Example: `(HQ+)+L?`.
- `Type` is the common message type that will be returned if the regex matches the message signature.

## TimeLocation
For internal use only. Should be ignored.

# Annotated structures
In order to read or write a message, an annotated structure is required. In a structure like this every field has to have an annotation. Fields without an annotation will be completely ignored. The annotation starts with the protocol name tag, then a colon, and some information enclosed in quotation marks.
``` go
type Structure struct {
    Field Type `astm:"information"`
}
```
The annotation name tag has to match the protocol we are using (also set in the configuration), and has to be all lower case. Valid options are:
``` go
astm
hl7
```

There are two distinct types of annotated structures from the [nesting layers](#nesting):
- Message or Group: This is the structure that represents multiple lines, or the entire message. Its fields consist of records or other groups.
- Record or Substructure: This is the structure that represents a single line in the message, or parts of it. Its fields are actual primitive data types or substructures.

## Key-value system
Inside the quotations after the protocol is the actual information, which is based on a key-value system. If there are more than one pair, they are listed with a separator character. Special attributes can optionally be added, which work the same way, nested with different separator characters.

### Syntax
At the top level the keys are always all capital letters, the list separator is a semicolon '**;**' and the part separator is an equal sign '**=**'.
``` go
KEY1=value1;KEY2=value2
```
At the attribute level the keys are always lowercase letters, the list separator is a comma '**,**' and the part separator is a colon '**:**'.
``` go
key1:value1,key2:value2
```

### Accepted keys
There is a predefined set of keys accepted in each role. The definitions can be seen in `constants/contstants.go`, but in the next chapter a detailed list and explanation is provided.

## Top level keys
| Key   | Value      | Where                  | On what | Explanation                                                         |
|-------|------------|------------------------|---------|---------------------------------------------------------------------|
| TAG   | tag        | Messages, Groups       | Record  | the identifying name tag of the record line (first field value)     |
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
The POS is mandatory on each field in a Record or Substructure. This is what positions the data elements in a record line. All the numbering start with 1. In Substructures, it is just a single integer, however in Records it is allowed to have multiple data fields pointing to the same Field, being components in it, basically a first level substructure not defined as separate type. In this case the field and component positions are separated by a dot.
Examples:
``` go
NoticeOfAdmissionCodestring    `hl7:"POS=24"`
NoticeOfAdmissionDatetime.Time `hl7:"POS=25"`

DataMeasurementValue     string `astm:"POS=4.1"`
InitialMeasurementValue  string `astm:"POS=4.2"`
MeasurementValueOfDevice string `astm:"POS=4.3"`
```
In ASTM the first two fields (position 1 and 2) are reserved for the record tag and the sequence number, so the actual data fields can only start from position 3. In HL7 only the first field (position 1) is reserved for the record tag.

### ATR
The ATR attributes have the value of a whole nested key-value list, as detailed below.

## Attribute level keys
| Key       | Value  | Where                   | On what       | Explanation                                                    |
|-----------|--------|-------------------------|---------------|----------------------------------------------------------------|
| optional  |        | Messages, Groups        | Record, Group | input record or section can be skipped if not found in message |
| subname   | name   | Messages, Groups        | Record        | additional identifier of the Record in the third field         |
| required  |        | Records, Substructures  | Field         | input is mandatory                                             |
| date      |        | Records, Substructures  | time.Time     | output as date only                                            |
| length    | number | Records, Substructures  | float         | output decimal length                                          |
| sequence  |        | Records, Substructures  | int           | marked as sequence number in HL7                               |

### optional
By default, at the message level all elements are mandatory, so the `optional` has to be added as an attribute. In case of a Group an optional attribute will recursively trickle down to all its elements becoming effectively optional even if it is not explicitly declared for them.
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
By default, at the record field level all fields are optional, so the `required` attribute has to be added to make it mandatory, and thus produce an error if the field is empty.

### date
The `date` attribute can be used on time.Time fields to signify that only the date part is relevant, and the time part should be ignored. In Marshal, it will be outputted in short date format (YYYYMMDD). It also has different behaviour with time zones as explained in the [Date and time](#date-and-time) section.
``` go
PlanEffectiveDate time.Time `hl7:"POS=13;ATR=date"
```

### length
The `length` attribute can be used on float fields to set the number of decimals to output. Detailed explanation is provided in the [Floating point precision](#floating-point-precision) section.
``` go
MeasuredValue float32 `astm:"POS=5;ATR=length:2"`
```

### sequence
The `sequence` attribute can be used on `type int` fields to signify that this field is the sequence number of the record in HL7 messages. This will make the field automatically incremented in marshal, and checked for validity in unmarshal if [enabled](#enforcesequencenumbercheck). ASTM has a fixed position (field 2) for the sequence number, so this attribute is not relevant there.
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
    Field3Component1 string `astm:"POS=3.1"`
    Field3Component2 string `astm:"POS=3.2"`
    Field3Component3 string `astm:"POS=3.3"`
}
```
```
R|1|comp1^comp2^comp3
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
R|1|comp1^comp2
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
Here are a few messages, groups, and arrays of them with annotations and their possible corresponding message lines.

### Simple message
``` go
type Lis02a2Message struct {
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
type HL7v23Message struct {
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
type GroupType struct {
    Record1 FirstRecordType  `hl7:"RT1"`
    Record2 SecontRecordType `hl7:"RT2"`
}
type HL7v23Message struct {
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
type Lis02a2Message struct {
    MessageHeader lis02a2.Header     `astm:"TAG=H"`
    GroupArray    []GroupType        `astm:"GROUP"`
    Terminator    lis02a2.Terminator `astm:"TAG=L"`
}
```
```
H|\^&||||
F|1|20251107|value2
S|1|value1|42
F|2|20251107|value2
S|1|value1|42
F|3|20251107|value2
S|1|value1|42
L|1|N
```
Note that the sequence number is incremented for each instance of the nested structure, however, only the first record of the nested structure takes the sequence number, and the rest is 1 (unless the nested structure has its own array inside).
