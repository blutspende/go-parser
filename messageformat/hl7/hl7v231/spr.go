package hl7v231

// SPR - Stored procedure request definition segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/SPR
type SPR struct {
	QueryTag string `hl7:"POS=2"`
	QueryResponseFormatCode string `hl7:"POS=3"`
	StoredProcedureName CE `hl7:"POS=4"`
	InputParameterList []QIP `hl7:"POS=5"`
}

