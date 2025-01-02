package wallet

type Constant struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type ConstantValueType string

const (
	ConstantValueU8     ConstantValueType = "u8"
	ConstantValueU16    ConstantValueType = "u16"
	ConstantValueU32    ConstantValueType = "u32"
	ConstantValueU64    ConstantValueType = "u64"
	ConstantValueString ConstantValueType = "string"
	ConstantValueBool   ConstantValueType = "boolean"
	ConstantValueRange  ConstantValueType = "range"
	ConstantValueBlob   ConstantValueType = "blob"
)

func ConstantDefaultU8(value uint8) Constant {
	return ConstantDefault(ConstantValueU8, value)
}

func ConstantDefaultU16(value uint16) Constant {
	return ConstantDefault(ConstantValueU16, value)
}

func ConstantDefaultU32(value uint32) Constant {
	return ConstantDefault(ConstantValueU32, value)
}

func ConstantDefaultU64(value uint64) Constant {
	return ConstantDefault(ConstantValueU64, value)
}

func ConstantDefaultString(value string) Constant {
	return ConstantDefault(ConstantValueString, value)
}

func ConstantDefaultBool(value string) Constant {
	return ConstantDefault(ConstantValueBool, value)
}

func ConstantDefaultBlob(value []byte) Constant {
	return ConstantDefault(ConstantValueBlob, value)
}

func ConstantDefault(value_type ConstantValueType, value interface{}) Constant {
	return Constant{
		Type: "default",
		Value: map[string]interface{}{
			"type":  value_type,
			"value": value,
		},
	}
}

func ConstantArray(value []Constant) Constant {
	return Constant{
		Type:  "array",
		Value: value,
	}
}
