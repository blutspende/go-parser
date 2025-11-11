package functions

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/models"
	"github.com/blutspende/go-parser/pconfig"
)

func ParseLine(inputLine string, targetStruct interface{}, recordAnnotation models.StructAnnotation, sequenceNumber int, config *pconfig.Configuration) (err error) {
	// Check for input line length
	if len(inputLine) == 0 {
		return errmsg.ErrLineParsingEmptyInput
	}

	// Setup inputFields variable to split the inputLine into and split start for header processing
	var inputFields []string
	splitStart := 0

	// Handle header special case
	isHeader := false
	if config.Protocol == pconfig.ASTM && inputLine[0] == 'H' {
		// Check if the inputLine is long enough to contain delimiters
		if len(inputLine) < 5 {
			return errmsg.ErrLineParsingHeaderTooShort
		}
		// Override delimiters
		config.Delimiters.Field = string(inputLine[1])
		config.Delimiters.Repeat = string(inputLine[2])
		config.Delimiters.Component = string(inputLine[3])
		config.Delimiters.Escape = string(inputLine[4])
		// Place the fix segment into the inputFields
		inputFields = []string{inputLine[0:1], inputLine[1:5]}
		splitStart = 6
		isHeader = true
	} else if config.Protocol == pconfig.HL7 && inputLine[0:3] == "MSH" {
		// Check if the inputLine is long enough to contain delimiters
		if len(inputLine) < 8 {
			return errmsg.ErrLineParsingHeaderTooShort
		}
		config.Delimiters.Field = inputLine[3:4]        // Default |
		config.Delimiters.Component = inputLine[4:5]    // Default ^
		config.Delimiters.Repeat = inputLine[5:6]       // Default ~
		config.Delimiters.Escape = inputLine[6:7]       // Default \
		config.Delimiters.SubComponent = inputLine[7:8] // Default &
		// Place the fix segment into the inputFields
		inputFields = []string{inputLine[0:3], inputLine[3:8]}
		splitStart = 9
		isHeader = true
	}

	// Split the input with the field delimiter starting from the split start position
	if len(inputLine) > splitStart {
		inputFields = append(inputFields, splitStringWithEscape(inputLine[splitStart:], config.Delimiters.Field, config)...)
	}

	// First two fields are mandatory in ASTM
	if config.Protocol == pconfig.ASTM && len(inputFields) < 2 {
		return fmt.Errorf("%w: %s", errmsg.ErrLineParsingNotEnoughFields, inputLine)
	}

	// Check for mach of name and subname
	// Note: name checking is always enforced, but instead of error it is returned in the nameOk variable
	if inputFields[0] != recordAnnotation.Tag {
		return fmt.Errorf("%w: expected '%s', got '%s', on line '%s'", errmsg.ErrLineParsingLineTagMismatch, recordAnnotation.Tag, inputFields[0], inputLine)
	}
	if subname, exists := recordAnnotation.Attributes[constants.AttributeSubname]; exists {
		// If subname is given at least 3 fields are required
		if len(inputFields) < 3 {
			return fmt.Errorf("%w: %s", errmsg.ErrLineParsingNotEnoughFields, inputLine)
		}
		// Check for subname match
		if inputFields[2] != subname {
			return fmt.Errorf("%w for subname: expected '%s', got '%s', on line '%s'", errmsg.ErrLineParsingLineTagMismatch, subname, inputFields[2], inputLine)
		}
	}

	// Check for validity of the sequence number if enforced - ASTM case
	if config.EnforceSequenceNumberCheck && config.Protocol == pconfig.ASTM && inputFields[1] != strconv.Itoa(sequenceNumber) && !isHeader {
		return fmt.Errorf("%w: expected '%s', got '%s', on line '%s'", errmsg.ErrLineParsingSequenceNumberMismatch, strconv.Itoa(sequenceNumber), inputFields[1], inputLine)
	}

	// Process the target structure
	targetTypes, targetValues, _, err := ProcessStructReflection(targetStruct)
	if err != nil {
		return err
	}

	// Iterate over the inputFields of the targetStruct struct
	for i, targetType := range targetTypes {
		// Parse the targetStruct field targetFieldAnnotation
		targetFieldAnnotation, err := ParseFieldAnnotation(targetType, config)
		if err != nil {
			if errors.Is(err, errmsg.ErrAnnotationParsingMissingAnnotation) {
				// If the annotation is missing, skip this field
				continue
			} else {
				return err
			}
		}

		// ASTM reserved: 1-name, 2-sequence number; HL7 reserved: 1-name
		if (config.Protocol == pconfig.ASTM && targetFieldAnnotation.FieldPos < 3) || (config.Protocol == pconfig.HL7 && targetFieldAnnotation.FieldPos < 2) {
			return errmsg.ErrLineParsingReservedFieldPosReference
		}

		// Check for validity of the sequence number if enforced - HL7 case
		_, sequenceAnnotation := targetFieldAnnotation.Attributes[constants.AttributeSequence]
		if config.EnforceSequenceNumberCheck && config.Protocol == pconfig.HL7 && sequenceAnnotation && inputFields[targetFieldAnnotation.FieldPos-1] != strconv.Itoa(sequenceNumber) {
			return fmt.Errorf("%w: expected '%s', got '%s', on line '%s'", errmsg.ErrLineParsingSequenceNumberMismatch, strconv.Itoa(sequenceNumber), inputFields[targetFieldAnnotation.FieldPos-1], inputLine)
		}

		// Not enough inputFields or empty inputField
		if len(inputFields) < targetFieldAnnotation.FieldPos || inputFields[targetFieldAnnotation.FieldPos-1] == "" {
			// If the field is required it's an error, otherwise skip it
			if _, exists := targetFieldAnnotation.Attributes[constants.AttributeRequired]; exists {
				return errmsg.ErrLineParsingRequiredInputFieldMissing
			} else {
				continue
			}
		}
		// Save the current inputField
		inputField := inputFields[targetFieldAnnotation.FieldPos-1]

		if targetFieldAnnotation.IsArray {
			// |rep1\rep2\rep3|
			// Field is an array
			repeats := splitStringWithEscape(inputField, config.Delimiters.Repeat, config)
			arrayType := reflect.SliceOf(targetValues[i].Type().Elem())
			arrayValue := reflect.MakeSlice(arrayType, len(repeats), len(repeats))
			for j, repeat := range repeats {
				if targetFieldAnnotation.IsSubstructure {
					// |comp1^comp2^comp3\comp1^comp2^comp3\comp1^comp2^comp3|
					// Substructures (with components) in the array: use parseSubstructure
					err = parseSubstructure(repeat, arrayValue.Index(j).Addr().Interface(), 1, config)
					if err != nil {
						return err
					}
				} else {
					// |value1\value2\value3|
					// Simple values in the array
					err = setField(repeat, arrayValue.Index(j), targetFieldAnnotation, config)
					if err != nil {
						return err
					}
				}

			}
			targetValues[i].Set(arrayValue)
		} else if targetFieldAnnotation.IsComponent {
			// |comp1^comp2^comp3|
			// Field is a component
			components := splitStringWithEscape(inputField, config.Delimiters.Component, config)
			// Not enough components in the inputField
			if len(components) < targetFieldAnnotation.ComponentPos {
				// Error if the component is required, skip otherwise
				if _, exists := targetFieldAnnotation.Attributes[constants.AttributeRequired]; exists {
					return errmsg.ErrLineParsingInputComponentsMissing
				} else {
					continue
				}
			}
			err = setField(components[targetFieldAnnotation.ComponentPos-1], targetValues[i], targetFieldAnnotation, config)
			if err != nil {
				return err
			}
		} else if targetFieldAnnotation.IsSubstructure {
			// |comp1^comp2^comp3|
			// If the field is a substructure use parseSubstructure to process it
			err = parseSubstructure(inputField, targetValues[i].Addr().Interface(), 1, config)
			if err != nil {
				return err
			}
		} else {
			// |field|
			// Field is not an array or component (normal singular field)
			err = setField(inputField, targetValues[i], targetFieldAnnotation, config)
			if err != nil {
				return err
			}
		}
		// Note: this could be a place to produce warnings about lost data
		// if i == targetFieldCount-1 && len(inputFields) > targetFieldAnnotation.FieldPos
	}
	// Return no error if everything went well
	return nil
}

func parseSubstructure(inputString string, targetStruct interface{}, depth int, config *pconfig.Configuration) (err error) {
	// Check depth limits
	if (config.Protocol == pconfig.ASTM && depth > 1) || (config.Protocol == pconfig.HL7 && depth > 2) {
		return errmsg.ErrLineParsingMaximumRecursionDepthExceeded
	}

	// Determine depth dependent delimiter
	delimiter := ""
	switch depth {
	case 1:
		delimiter = config.Delimiters.Component
	case 2:
		delimiter = config.Delimiters.SubComponent
	default:
		return errmsg.ErrLineParsingInvalidRecursionDepth
	}
	// Split the input with the field delimiter
	inputFields := splitStringWithEscape(inputString, delimiter, config)

	// Process the target structure
	targetTypes, targetValues, _, err := ProcessStructReflection(targetStruct)
	if err != nil {
		return err
	}

	// Iterate over the inputFields of the targetStruct struct
	for i, targetType := range targetTypes {
		// Parse the targetStruct field targetFieldAnnotation
		targetFieldAnnotation, err := ParseFieldAnnotation(targetType, config)
		if err != nil {
			if errors.Is(err, errmsg.ErrAnnotationParsingMissingAnnotation) {
				// If the annotation is missing, skip this field
				continue
			} else {
				return err
			}
		}

		// Not enough inputFields or empty inputField
		if len(inputFields) < targetFieldAnnotation.FieldPos || inputFields[targetFieldAnnotation.FieldPos-1] == "" {
			// If the field is required it's an error, otherwise skip it
			if _, exists := targetFieldAnnotation.Attributes[constants.AttributeRequired]; exists {
				return errmsg.ErrLineParsingRequiredInputFieldMissing
			} else {
				continue
			}
		}
		// Save the current inputField
		inputField := inputFields[targetFieldAnnotation.FieldPos-1]

		// Recurse if the target is also a struct
		if targetValues[i].Kind() == reflect.Struct && targetFieldAnnotation.IsSubstructure {
			err = parseSubstructure(inputField, targetValues[i].Addr().Interface(), depth+1, config)
		} else {
			// Field is single value: set it
			err = setField(inputField, targetValues[i], targetFieldAnnotation, config)
		}
		if err != nil {
			return err
		}
	}

	// Return no error if everything went well
	return nil
}

func setField(value string, field reflect.Value, annotation models.FieldAnnotation, config *pconfig.Configuration) (err error) {
	// Ensure the field is settable
	if !field.CanSet() {
		// Field is not settable
		return errmsg.ErrLineParsingNonSettableField
	}
	// Handle pointer types by allocating and dereferencing
	if field.Kind() == reflect.Ptr {
		elemType := field.Type().Elem()
		newElem := reflect.New(elemType)
		field.Set(newElem)
		field = field.Elem()
	}
	// Set the field value
	switch field.Kind() {
	case reflect.String:
		escaped := value
		if config.Protocol == pconfig.ASTM {
			escaped = filterStringEscapeChars(value, config.Delimiters.Escape)
		} else if config.Protocol == pconfig.HL7 {
			escaped, err = replaceHL7Escapes(value, config)
			if err != nil {
				return err
			}
		}
		if field.Type().ConvertibleTo(reflect.TypeOf("")) {
			field.Set(reflect.ValueOf(escaped).Convert(field.Type()))
		} else {
			field.Set(reflect.ValueOf(escaped))
		}
		return nil
	case reflect.Int:
		num, err := strconv.Atoi(value)
		if err != nil {
			return errmsg.ErrLineParsingDataParsingError
		}
		field.Set(reflect.ValueOf(num))
		return nil
	case reflect.Float32:
		num, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return errmsg.ErrLineParsingDataParsingError
		}
		field.Set(reflect.ValueOf(float32(num)))
		return nil
	case reflect.Float64:
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errmsg.ErrLineParsingDataParsingError
		}
		field.Set(reflect.ValueOf(num))
		return nil
	// Check for time.Time type (it reflects as a Struct)
	case reflect.Struct:
		if field.Type() == reflect.TypeOf(time.Time{}) {
			var timeFormat string
			var isShort bool
			switch len(value) {
			case 8:
				timeFormat = "20060102" // YYYYMMDD
				isShort = true
			case 14:
				timeFormat = "20060102150405" // YYYYMMDDHHMMSS
				isShort = false
			default:
				return errmsg.ErrLineParsingInvalidDateFormat
			}
			timeInLocation, err := time.ParseInLocation(timeFormat, value, config.TimeLocation)
			if err != nil {
				return errmsg.ErrLineParsingDataParsingError
			}
			// Keep the short format in local time (if configured so) to preserve the actual day
			if isShort && config.KeepShortDateTimeZone {
				timeInLocation = timeInLocation.In(config.TimeLocation)
			} else {
				// Convert the long format to UTC
				timeInLocation = timeInLocation.UTC()
			}
			field.Set(reflect.ValueOf(timeInLocation))
			return nil
		} else {
			// Note: option to handle other struct types here
		}
	}
	// Return error if no type match was found (each successful parsing returns nil)
	return errmsg.ErrLineParsingUnsupportedDataType
}

func splitStringWithEscape(input string, delimiter string, config *pconfig.Configuration) (result []string) {
	inputRunes := []rune(input)
	delimiterRune := rune(delimiter[0])
	if config.Protocol == pconfig.ASTM {
		escapeRune := rune(config.Delimiters.Escape[0])
		start := 0
		for i := 0; i < len(inputRunes); i++ {
			if inputRunes[i] == delimiterRune {
				result = append(result, string(inputRunes[start:i]))
				start = i + 1
			}
			if i == len(inputRunes)-1 {
				result = append(result, string(inputRunes[start:i+1]))
			}
			if inputRunes[i] == escapeRune {
				i++
				continue
			}
		}
	} else if config.Protocol == pconfig.HL7 {
		result = strings.Split(input, delimiter)
	}
	return result
}

func filterStringEscapeChars(input string, escape string) string {
	var builder strings.Builder
	escapeRune := rune(escape[0])
	inputRunes := []rune(input)
	for i := 0; i < len(inputRunes); i++ {
		if inputRunes[i] == escapeRune {
			i++
			if i < len(inputRunes) {
				builder.WriteRune(inputRunes[i])
			}
		} else {
			builder.WriteRune(inputRunes[i])
		}
	}
	return builder.String()
}

func replaceHL7Escapes(input string, config *pconfig.Configuration) (result string, err error) {
	var builder strings.Builder
	escapeRune := rune(config.Delimiters.Escape[0])
	inputRunes := []rune(input)
	for i := 0; i < len(inputRunes); i++ {
		if inputRunes[i] == escapeRune {
			var end int
			for end = i + 1; end < len(inputRunes) && inputRunes[end] != escapeRune; end++ {
			}
			if end == len(inputRunes) {
				return "", errmsg.ErrLineParsingUnterminatedEscapeSequence
			}
			escapeSeq := string(inputRunes[i+1 : end])
			switch escapeSeq {
			case "F":
				builder.WriteString(config.Delimiters.Field)
			case "S":
				builder.WriteString(config.Delimiters.Component)
			case "R":
				builder.WriteString(config.Delimiters.Repeat)
			case "E":
				builder.WriteString(config.Delimiters.Escape)
			case "T":
				builder.WriteString(config.Delimiters.SubComponent)
			case ".br":
				builder.WriteString("\r")
			default:
				if regexp.MustCompile("^X[0-9A-F]{2}$").MatchString(escapeSeq) {
					charNumStr := escapeSeq[1:]
					charInt, _ := strconv.ParseInt(charNumStr, 16, 32)
					builder.WriteRune(rune(charInt))
				} else {
					return "", errmsg.ErrLineParsingUnknownEscapeSequence
				}
			}
			i = end
		} else {
			builder.WriteRune(inputRunes[i])
		}
	}
	return builder.String(), nil
}
