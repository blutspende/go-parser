package functions

import (
	"errors"
	"fmt"

	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errdef"
	"github.com/blutspende/go-parser/pconfig"

	"reflect"
)

func ParseStruct(inputLines []string, targetStruct interface{}, config *pconfig.Configuration) (err error) {
	lineIndex := 0
	return parseStructRecursive(inputLines, targetStruct, &lineIndex, 1, false, 0, config)
}

func parseStructRecursive(inputLines []string, targetStruct interface{}, lineIndex *int, sequenceNumber int, optional bool, depth int, config *pconfig.Configuration) (err error) {
	// Unless at leat one matching element is found the structure is empty
	empty := true
	// Check for maximum depth
	if depth >= constants.MaxDepth {
		return errdef.ErrStructureParsingMaxDepthReached
	}
	// Process the target structure
	targetTypes, targetValues, _, err := ProcessStructReflection(targetStruct)
	if err != nil {
		return err
	}

	// Iterate over the inputFields of the targetStruct struct
	for i, targetType := range targetTypes {
		// Parse the targetStruct field targetFieldAnnotation
		targetStructAnnotation, err := ParseStructAnnotation(targetType, config)
		if err != nil {
			if errors.Is(err, errdef.ErrAnnotationParsingMissingAnnotation) {
				// If the annotation is missing, skip this field
				continue
			} else {
				return err
			}
		}
		// Save the target value pointer
		targetValue := targetValues[i].Addr().Interface()
		// Set optionality from target annotation and recursive parameter from above
		_, targetOptionalAttribute := targetStructAnnotation.Attributes[constants.AttributeOptional]
		targetOptional := targetOptionalAttribute || optional

		// Target is an array it is iterated with conditional break (unknown length)
		if targetStructAnnotation.IsArray {
			// Create the array structure
			sliceType := reflect.SliceOf(targetValues[i].Type().Elem())
			targetValues[i].Set(reflect.MakeSlice(sliceType, 0, 0))

			// Iterate as long as we have matching input structure and still have input lines
			for seq := 1; *lineIndex < len(inputLines); seq++ {
				// Create a new element for the slice to parse into
				elem := reflect.New(targetValues[i].Type().Elem()).Elem()

				if targetStructAnnotation.IsGroup {
					// Composite target: recursively parse the composite structure
					err = parseStructRecursive(inputLines, elem.Addr().Interface(), lineIndex, seq, targetOptional, depth+1, config)
				} else {
					// Non-composite target: parse the line into the new element
					err = ParseLine(inputLines[*lineIndex], elem.Addr().Interface(), targetStructAnnotation, seq, config)
					// If the line parsed successfully: increment line index
					if err == nil {
						*lineIndex++
					}
				}
				if err != nil {
					// Line tag mismatch or empty structure means the array is over
					if errors.Is(err, errdef.ErrLineParsingLineTagMismatch) || errors.Is(err, errdef.ErrStructureParsingEmptyStructure) {
						break
					} else {
						// Any other error is returned as error
						return err
					}
				}

				// If no error: add the new element to the slice and mark the struct as not empty
				targetValues[i].Set(reflect.Append(targetValues[i], elem))
				empty = false
			}
		} else {
			// Single element structure
			// Check for enough input lines
			if *lineIndex >= len(inputLines) {
				if targetOptional {
					// Target is optional: it can be skipped (to allow the loop to find any remaining non-optional targets)
					continue
				} else {
					if config.EnforceMessageCompleteness {
						// If not optional and completeness is enforced: error
						return fmt.Errorf("%w: expected type '%s'", errdef.ErrStructureParsingInputLinesDepleted, targetType.Name)
					} else {
						// Completeness is not enforced: just return (nothing more to parse)
						return nil
					}
				}
			}

			if targetStructAnnotation.IsGroup {
				// Composite target: go further down the rabbit hole
				err = parseStructRecursive(inputLines, targetValue, lineIndex, 1, targetOptional, depth+1, config)
			} else {
				// Non-composite target: there is a single line to parse
				// Determine sequence number: first element inherits from the parent call, the rest is 1
				seq := 1
				if i == 0 {
					seq = sequenceNumber
				}
				// Parse the line and increment the line index
				err = ParseLine(inputLines[*lineIndex], targetValue, targetStructAnnotation, seq, config)
				// If the line parsed successfully: increment line index
				if err == nil {
					*lineIndex++
				}
			}

			if err != nil {
				// If optional: line tag mismatch or empty structure can be skipped
				if targetOptional && (errors.Is(err, errdef.ErrLineParsingLineTagMismatch) || errors.Is(err, errdef.ErrStructureParsingEmptyStructure)) {
					continue
				} else {
					// Any other case: the error is returned as error
					return err
				}
			} else {
				// If parsed successfully: mark the structure as not empty
				empty = false
			}
		}
	}

	// Return the empty structure error if nothing matched
	if empty {
		return errdef.ErrStructureParsingEmptyStructure
	} else {
		// Or no error if all ok
		return nil
	}
}
