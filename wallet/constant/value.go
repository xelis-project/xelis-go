package constant

type ValueType string

const (
	ValueTypeNull   ValueType = "null"
	ValueTypeU8     ValueType = "u8"
	ValueTypeU16    ValueType = "u16"
	ValueTypeU32    ValueType = "u32"
	ValueTypeU64    ValueType = "u64"
	ValueTypeString ValueType = "string"
	ValueTypeBool   ValueType = "boolean"
	ValueTypeBlob   ValueType = "blob"
	ValueTypeRange  ValueType = "range"
)

type Value struct {
	Type  ValueType   `json:"type"`
	Value interface{} `json:"value,omitempty"`
}

func ValueNull() Value {
	return Value{Type: ValueTypeNull}
}

func ValueU8(value uint8) Value {
	return Value{
		Type:  ValueTypeU8,
		Value: value,
	}
}

func ValueU16(value uint16) Value {
	return Value{
		Type:  ValueTypeU16,
		Value: value,
	}
}

func ValueU32(value uint32) Value {
	return Value{
		Type:  ValueTypeU32,
		Value: value,
	}
}

func ValueU64(value uint64) Value {
	return Value{
		Type:  ValueTypeU64,
		Value: value,
	}
}

func ValueString(value string) Value {
	return Value{
		Type:  ValueTypeString,
		Value: value,
	}
}

func ValueBool(value bool) Value {
	return Value{
		Type:  ValueTypeBool,
		Value: value,
	}
}

func ValueBlob(value []uint) Value {
	return Value{
		Type:  ValueTypeBlob,
		Value: value,
	}
}

func ValueRange(start Value, end Value) Value {
	if start.Type != end.Type {
		panic("start / end are not the same type")
	}

	return Value{
		Type: ValueTypeRange,
		Value: []Value{
			start,
			end,
			{Type: start.Type},
		},
	}
}
