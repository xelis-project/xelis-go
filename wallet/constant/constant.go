package constant

import (
	"math/big"
	"reflect"
)

type ConstantType string

const (
	ConstantTypeDefault ConstantType = "default"

	ConstantTypeArray    ConstantType = "array"
	ConstantTypeOptional ConstantType = "optional"
	ConstantTypeStruct   ConstantType = "struct" // TODO
	ConstantTypeMap      ConstantType = "map"    // TODO
	ConstantTypeEnum     ConstantType = "enum"   // TODO
)

type Constant struct {
	Type  ConstantType `json:"type"`
	Value interface{}  `json:"value,omitempty"`
}

type Struct struct {
	Id     uint16        `json:"id"`
	Fields []interface{} `json:"fields"`
}

func Default(value Value) Constant {
	return Constant{
		Type:  ConstantTypeDefault,
		Value: value,
	}
}

func DefaultU8(value uint8) Constant {
	return Default(ValueU8(value))
}

func DefaultU16(value uint16) Constant {
	return Default(ValueU16(value))
}

func DefaultU32(value uint32) Constant {
	return Default(ValueU32(value))
}

func DefaultU64(value uint64) Constant {
	return Default(ValueU64(value))
}

func DefaultU128(value *big.Int) Constant {
	return Default(ValueU128(value))
}

func DefaultU256(value *big.Int) Constant {
	return Default(ValueU256(value))
}

func DefaultString(value string) Constant {
	return Default(ValueString(value))
}

func DefaultBlob(value []uint) Constant {
	return Default(ValueBlob(value))
}

func DefaultBool(value bool) Constant {
	return Default(ValueBool(value))
}

func DefaultRange(start interface{}, end interface{}) Constant {
	if reflect.TypeOf(start) != reflect.TypeOf(end) {
		panic("start / end are not the same type")
	}

	switch s := start.(type) {
	case uint8:
		return Default(ValueRange(ValueU8(s), ValueU8(end.(uint8))))
	case uint16:
		return Default(ValueRange(ValueU16(s), ValueU16(end.(uint16))))
	case uint32:
		return Default(ValueRange(ValueU32(s), ValueU32(end.(uint32))))
	case uint64:
		return Default(ValueRange(ValueU64(s), ValueU64(end.(uint64))))
	}

	// makes to sense to create range for bool, string & blob
	panic("not supported type")
}

func Array(value []Constant) Constant {
	return Constant{
		Type:  ConstantTypeArray,
		Value: value,
	}
}

func Optional(value Constant) Constant {
	return Constant{
		Type:  ConstantTypeOptional,
		Value: value,
	}
}
