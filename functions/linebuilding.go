package functions

import (
	"errors"
	"math"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/blutspende/go-parser/constants"
	notationconst "github.com/blutspende/go-parser/enums/notation"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/models"
	"github.com/blutspende/go-parser/pconfig"
)

func BuildLine(sourceStruct interface{}, lineTypeName string, sequenceNumber int, config *pconfig.Configuration) (result string, err error) {
	// Process the target structure
	sourceTypes, sourceValues, sourceTypesLength, err := ProcessStructReflection(sourceStruct)
	if err != nil {
		return "", err
	}

	// Create a map to store field values indexed by FieldPos
	fieldMap := make(map[int]string)

	// Create an array to store already processed component fields to allow any sparse placement
	processedComponentFields := make([]int, 0)

	// Handle fixed values (line name, header delimiters, ASTM auto sequence)
	fieldMap[1] = lineTypeName
	if config.Protocol == pconfig.ASTM {
		if lineTypeName == "H" {
			fieldMap[2] = config.Delimiters.Repeat +
				config.Delimiters.Component +
				config.Delimiters.Escape
		} else {
			fieldMap[2] = strconv.Itoa(sequenceNumber)
		}
	}
	if config.Protocol == pconfig.HL7 && lineTypeName == "MSH" {
		fieldMap[2] = config.Delimiters.Component +
			config.Delimiters.Repeat +
			config.Delimiters.Escape +
			config.Delimiters.SubComponent
	}

	// Iterate over the inputFields of the targetStruct struct
	for i := 0; i < sourceTypesLength; i++ {
		// Parse the sourceStruct field sourceFieldAnnotation
		sourceFieldAnnotation, err := ParseFieldAnnotation(sourceTypes[i], config)
		if err != nil {
			if errors.Is(err, errmsg.ErrAnnotationParsingMissingAnnotation) {
				// If the annotation is missing, skip this field
				continue
			} else {
				return "", err
			}
		}

		// ASTM reserved: 1-name, 2-sequence number; HL7 reserved: 1-name
		if (config.Protocol == pconfig.ASTM && sourceFieldAnnotation.FieldPos < 3) || (config.Protocol == pconfig.HL7 && sourceFieldAnnotation.FieldPos < 2) {
			return "", errmsg.ErrLineBuildingReservedFieldPosReference
		}

		// Copy the sequence number into the field with the sequence annotation - HL7 case
		_, sequenceAnnotation := sourceFieldAnnotation.Attributes[constants.AttributeSequence]
		if config.Protocol == pconfig.HL7 && sequenceAnnotation {
			sourceValues[i] = reflect.ValueOf(sequenceNumber)
		}

		fieldValueString := ""
		// If the field is an array, iterate over its elements and use the Repeat delimiter
		if sourceFieldAnnotation.IsArray {
			for j := 0; j < sourceValues[i].Len(); j++ {
				elementValue := sourceValues[i].Index(j)
				convertedValue := ""
				if sourceFieldAnnotation.IsSubstructure {
					// If the field is a substructure use buildSubstructure to process it
					convertedValue, err = buildSubstructure(elementValue.Interface(), 1, config)
					if err != nil {
						return "", err
					}
				} else {
					// Simple field, convert it directly
					convertedValue, err = convertField(elementValue, sourceFieldAnnotation, config)
					if err != nil {
						return "", err
					}
				}
				fieldValueString += convertedValue
				if j < sourceValues[i].Len()-1 {
					fieldValueString += config.Delimiters.Repeat
				}
			}
		} else if sourceFieldAnnotation.IsComponent {
			if slices.Contains(processedComponentFields, sourceFieldAnnotation.FieldPos) {
				// If the field is already processed, skip it
				continue
			}
			// Create a map to store the component values indexed by ComponentPos
			componentMap := make(map[int]string)
			// Iterate over the whole inputFields of the targetStruct struct to find the components anywhere
			for j := 0; j < sourceTypesLength; j++ {
				// Parse the targetStruct field targetFieldAnnotation
				currentFieldAnnotation, err := ParseFieldAnnotation(sourceTypes[j], config)
				if err != nil {
					if errors.Is(err, errmsg.ErrAnnotationParsingMissingAnnotation) {
						// If the annotation is missing, skip this field
						continue
					} else {
						return "", err
					}
				}
				// If the field number is the same as the sourceFieldAnnotation, process it
				if currentFieldAnnotation.FieldPos == sourceFieldAnnotation.FieldPos {
					// Convert current component
					componentValue, err := convertField(sourceValues[j], currentFieldAnnotation, config)
					if err != nil {
						return "", err
					}
					// Store the value in the component map
					componentMap[currentFieldAnnotation.ComponentPos] = componentValue
				}
			}
			// Construct the result into the fieldValueString
			fieldValueString = constructResult(componentMap, config.Delimiters.Component, config.Notation)
			// Mark the field as processed
			processedComponentFields = append(processedComponentFields, sourceFieldAnnotation.FieldPos)
		} else if sourceFieldAnnotation.IsSubstructure {
			// If the field is a substructure use buildSubstructure to process it
			fieldValueString, err = buildSubstructure(sourceValues[i].Interface(), 1, config)
			if err != nil {
				return "", err
			}
		} else {
			// If the field is not an array, convert it directly
			fieldValueString, err = convertField(sourceValues[i], sourceFieldAnnotation, config)
			if err != nil {
				return "", err
			}
		}

		// Store the field value in the map using FieldPos as the key
		fieldMap[sourceFieldAnnotation.FieldPos] = fieldValueString
	}

	// Construct the result string based on the field map
	result = constructResult(fieldMap, config.Delimiters.Field, config.Notation)

	// Handle empty record dropping (empty lines will be filtered out later)
	if config.DropEmptyOutputRecords && result == lineTypeName {
		return "", nil
	}
	// Handle HL7 empty line mandatory field separator
	if config.Protocol == pconfig.HL7 && result == lineTypeName {
		result += config.Delimiters.Field
	}
	// Return the result with no error
	return result, nil
}

func buildSubstructure(sourceStruct interface{}, depth int, config *pconfig.Configuration) (result string, err error) {
	// Check depth limits
	if (config.Protocol == pconfig.ASTM && depth > 1) || (config.Protocol == pconfig.HL7 && depth > 2) {
		return "", errmsg.ErrLineBuildingMaximumRecursionDepthExceeded
	}

	// Process the target structure
	sourceTypes, sourceValues, sourceTypesLength, err := ProcessStructReflection(sourceStruct)
	if err != nil {
		return "", err
	}

	// Create a map to store component values indexed by FieldPos
	componentMap := make(map[int]string)

	// Iterate over the inputFields of the targetStruct struct
	for i := 0; i < sourceTypesLength; i++ {
		// Parse the sourceStruct field sourceFieldAnnotation
		sourceFieldAnnotation, err := ParseFieldAnnotation(sourceTypes[i], config)
		if err != nil {
			if errors.Is(err, errmsg.ErrAnnotationParsingMissingAnnotation) {
				// If the annotation is missing, skip this field
				continue
			} else {
				return "", err
			}
		}

		// Recurse if the source is also a struct
		var componentValueString string
		if sourceValues[i].Kind() == reflect.Struct && sourceFieldAnnotation.IsSubstructure {
			componentValueString, err = buildSubstructure(sourceValues[i].Addr().Interface(), depth+1, config)
		} else {
			// Convert the component directly
			componentValueString, err = convertField(sourceValues[i], sourceFieldAnnotation, config)
		}
		if err != nil {
			return "", err
		}

		// Store the component value in the map using FieldPos as the key
		componentMap[sourceFieldAnnotation.FieldPos] = componentValueString
	}

	// Determine depth dependent delimiter
	delimiter := ""
	switch depth {
	case 1:
		delimiter = config.Delimiters.Component
	case 2:
		delimiter = config.Delimiters.SubComponent
	default:
		return "", errmsg.ErrLineBuildingInvalidRecursionDepth
	}

	// Construct the result string
	result = constructResult(componentMap, delimiter, config.Notation)

	// Return result with no error
	return result, nil
}

func constructResult(fieldMap map[int]string, delimiter string, notation string) (result string) {
	// Determine how many fields to include by finding the biggest index
	lastIndex := 0
	for key := range fieldMap {
		// In short notation only non-empty fields are included at the end
		if key > lastIndex && (!(notation == notationconst.Short) || fieldMap[key] != "") {
			lastIndex = key
		}
	}
	// Iterate from one to the last index, building the result string
	result = ""
	for i := 1; i <= lastIndex; i++ {
		// If the field exists in the map, append its value to the result
		if value, exists := fieldMap[i]; exists {
			result += value
		}
		// Add the delimiter after all but the last element
		if i != lastIndex {
			result += delimiter
		}
	}
	// Return the built string
	return result
}

func convertField(field reflect.Value, annotation models.FieldAnnotation, config *pconfig.Configuration) (result string, err error) {
	// Check if the field is a pointer, nil returns empty, otherwise dereference it
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			return "", nil
		}
		field = field.Elem()
	}
	// Format the result as a string based on the field type
	switch field.Kind() {
	case reflect.String:
		if field.Type().ConvertibleTo(reflect.TypeOf("")) {
			if config.EscapeOutputStrings {
				result = buildStringEscapeChars(field.String(), config)
			} else {
				result = field.String()
			}
		} else {
			return "", errmsg.ErrLineBuildingUsupportedDataType
		}
		return result, nil
	case reflect.Int:
		result = strconv.Itoa(int(field.Int()))
		return result, nil
	case reflect.Float32, reflect.Float64:
		precision := config.DefaultDecimalPrecision
		if value, exists := annotation.Attributes[constants.AttributeLength]; exists {
			precision, err = strconv.Atoi(value)
			if err != nil {
				return "", errmsg.ErrLineBuildingInvalidLengthAttributeValue
			}
		}
		result = strconv.FormatFloat(field.Float(), 'f', precision, field.Type().Bits())
		if !config.RoundLastDecimal && precision >= 0 {
			factor := math.Pow(10, float64(precision))
			truncated := math.Trunc(field.Float()*factor) / factor
			result = strconv.FormatFloat(truncated, 'f', precision, field.Type().Bits())
		}
		return result, nil
	case reflect.Struct:
		// Check for time.Time type (it reflects as a Struct)
		if field.Type() == reflect.TypeOf(time.Time{}) {
			timeFormat := "20060102150405"
			if _, exists := annotation.Attributes[constants.AttributeDate]; exists {
				timeFormat = "20060102"
			}
			// Check if the field is a time.Time
			timeValue, ok := field.Interface().(time.Time)
			if !ok {
				return "", errmsg.ErrLineBuildingInvalidDateFormat
			}
			// Return empty if the time is zero
			if timeValue.IsZero() {
				return "", nil
			}
			// Convert the time to the config's timezone
			timeInLocation := timeValue.In(config.TimeLocation)
			// Format the date as a string
			result = timeInLocation.Format(timeFormat)
			return result, nil
		} else {
			// Note: option to handle other struct types here
		}
	}
	// Return error if no type match was found (each successful conversion returns with nil)
	return "", errmsg.ErrLineBuildingUsupportedDataType
}

func buildStringEscapeChars(input string, config *pconfig.Configuration) string {
	var builder strings.Builder
	inputRunes := []rune(input)
	esc := string(config.Delimiters.Escape[0])
	if config.Protocol == pconfig.ASTM {
		for i := 0; i < len(inputRunes); i++ {
			if inputRunes[i] == rune(config.Delimiters.Field[0]) ||
				inputRunes[i] == rune(config.Delimiters.Repeat[0]) ||
				inputRunes[i] == rune(config.Delimiters.Component[0]) ||
				inputRunes[i] == rune(config.Delimiters.Escape[0]) {
				builder.WriteString(esc)
			}
			builder.WriteRune(inputRunes[i])
		}
	} else if config.Protocol == pconfig.HL7 {
		for i := 0; i < len(inputRunes); i++ {
			switch inputRunes[i] {
			case rune(config.Delimiters.Field[0]):
				builder.WriteString(esc + "F" + esc)
			case rune(config.Delimiters.Repeat[0]):
				builder.WriteString(esc + "R" + esc)
			case rune(config.Delimiters.Component[0]):
				builder.WriteString(esc + "S" + esc)
			case rune(config.Delimiters.SubComponent[0]):
				builder.WriteString(esc + "T" + esc)
			case rune(config.Delimiters.Escape[0]):
				builder.WriteString(esc + "E" + esc)
			case '\r':
				builder.WriteString(esc + "X0D" + esc)
			default:
				builder.WriteRune(inputRunes[i])
			}
		}
	} else {
		// Note: this is an error case, but not handled expecting prior protocol verification
		return ""
	}
	return builder.String()
}
