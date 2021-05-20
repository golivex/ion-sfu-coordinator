// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsdata

type DecimalReturnType string

// Enum values for DecimalReturnType
const (
	DecimalReturnTypeDoubleOrLong DecimalReturnType = "DOUBLE_OR_LONG"
	DecimalReturnTypeString       DecimalReturnType = "STRING"
)

func (enum DecimalReturnType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DecimalReturnType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TypeHint string

// Enum values for TypeHint
const (
	TypeHintDate      TypeHint = "DATE"
	TypeHintDecimal   TypeHint = "DECIMAL"
	TypeHintTime      TypeHint = "TIME"
	TypeHintTimestamp TypeHint = "TIMESTAMP"
)

func (enum TypeHint) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TypeHint) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
