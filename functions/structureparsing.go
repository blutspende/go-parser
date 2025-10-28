package functions

import (
	"fmt"

	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/parserconfig"

	"reflect"
)

func ParseStruct(inputLines []string, targetStruct interface{}, config *parserconfig.Configuration) (err error) {
	lineIndex := 0
	_, err = parseStructRecursive(inputLines, targetStruct, &lineIndex, 1, false, 0, config)
	return
}

func parseStructRecursive(inputLines []string, targetStruct interface{}, lineIndex *int, sequenceNumber int, optional bool, depth int, config *parserconfig.Configuration) (structMatch bool, err error) {
	// Assume the whole structure is a mismatches until at leat one matching element is found
	structMatch = false
	// Check for maximum depth
	if depth >= constants.MaxDepth {
		return structMatch, errmsg.ErrStructureParsingMaxDepthReached
	}
	// Process the target structure
	targetTypes, targetValues, _, err := ProcessStructReflection(targetStruct)
	if err != nil {
		return structMatch, err
	}

	// Iterate over the inputFields of the targetStruct struct
	for i, targetType := range targetTypes {
		// Parse the targetStruct field targetFieldAnnotation
		targetStructAnnotation, err := ParseStructAnnotation(targetType, config)
		if err != nil {
			return structMatch, err
		}
		// Save the target value pointer
		targetValue := targetValues[i].Addr().Interface()
		// Set optionality from target annotation and recursive parameter from above
		_, targetOptionalAttribute := targetStructAnnotation.Attributes[constants.AttributeOptional]
		targetOptional := targetOptionalAttribute || optional
		// Note: arrays are implicitly optional (empty arrays are valid)
		targetOptional = targetOptional || targetStructAnnotation.IsArray

		// Target is an array it is iterated with conditional break (unknown length)
		if targetStructAnnotation.IsArray {
			// Create the array structure
			sliceType := reflect.SliceOf(targetValues[i].Type().Elem())
			targetValues[i].Set(reflect.MakeSlice(sliceType, 0, 0))
			// Iterate as long as we have matching input structure and still have input lines
			for seq := 1; *lineIndex < len(inputLines); seq++ {
				// Create a new element for the slice to parse into
				elem := reflect.New(targetValues[i].Type().Elem()).Elem()
				// Match value communicates if the line or substructure can be found at the index in the input lines
				var localMatch bool
				if targetStructAnnotation.IsGroup {
					// Composite target: recursively parse the composite structure
					localMatch, err = parseStructRecursive(inputLines, elem.Addr().Interface(), lineIndex, seq, targetOptional, depth+1, config)
				} else {
					// Non-composite target: parse the line into the new element
					localMatch, err = ParseLine(inputLines[*lineIndex], elem.Addr().Interface(), targetStructAnnotation, seq, config)
					// If the line matched: increment line index
					if localMatch {
						*lineIndex++
					}
				}
				// Handle potential error from both cases
				if err != nil {
					return structMatch, err
				}
				// Element mismatch means the end of the array
				if !localMatch {
					break
				} else {
					// Element matched: the structure matches
					structMatch = true
				}

				// If matched and no error: add the new element to the slice
				targetValues[i].Set(reflect.Append(targetValues[i], elem))
			}
		} else {
			// Single element structure
			// Check for enough input lines
			if *lineIndex >= len(inputLines) {
				if targetOptional {
					// Target is optional: it can be skipped (to allow the loop to find any non-optional remaining)
					continue
				} else {
					if config.EnforceMessageCompleteness {
						// If not optional and completeness is enforced: error
						return structMatch, fmt.Errorf("%w: expected [%s]", errmsg.ErrStructureParsingInputLinesDepleted, targetStructAnnotation.Tag)
					} else {
						// Completeness is not enforced: just return (nothing more to parse)
						return structMatch, nil
					}
				}
			}
			var localMatch bool
			if targetStructAnnotation.IsGroup {
				// Composite target: go further down the rabbit hole
				localMatch, err = parseStructRecursive(inputLines, targetValue, lineIndex, 1, targetOptional, depth+1, config)
				if err != nil {
					return structMatch, err
				}
			} else {
				// Non-composite target: there is a single line to parse
				// Determine sequence number: first element inherits from the parent call, the rest is 1
				seq := 1
				if i == 0 {
					seq = sequenceNumber
				}
				// Parse the line and increment the line index
				localMatch, err = ParseLine(inputLines[*lineIndex], targetValue, targetStructAnnotation, seq, config)
				if err != nil {
					return structMatch, err
				}
				// If the line matched: increment line index
				if localMatch {
					*lineIndex++
				}
			}
			// Element mismatch
			if !localMatch {
				if targetOptional {
					// Target is optional: it can be skipped
					continue
				} else {
					// Target not optional: return with mismatch error
					return structMatch, fmt.Errorf("%w: @ln %d, expected [%s]", errmsg.ErrStructureParsingStructureMismatch, *lineIndex, targetStructAnnotation.Tag)
				}
			} else {
				// Element matched: the structure matches
				structMatch = true
			}
		}
	}
	// Return the match state and no error
	return structMatch, nil
}
