package extra_data

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

type DataValueReader struct {
	Reader *bytes.Reader
}

func (d *DataValueReader) Read() (dataElement Element, err error) {
	dataElementType, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	switch dataElementType {
	case byte(ElementValue): // Value
		var value Value
		value, err = d.readDataValue()
		if err != nil {
			return
		}

		dataElement = Element{Value: value}
	case byte(ElementArray): // Array
		var size byte
		size, err = d.Reader.ReadByte()
		if err != nil {
			return
		}

		var values []Element
		for i := 0; i < int(size); i++ {
			var value Element
			value, err = d.Read()
			if err != nil {
				return
			}

			values = append(values, value)
		}

		dataElement = Element{Array: values}
	case byte(ElementFields): // Fields / Map
		var size byte
		size, err = d.Reader.ReadByte()
		if err != nil {
			return
		}

		fields := make(map[Value]Element, 0)
		for i := 0; i < int(size); i++ {
			var key Value
			key, err = d.readDataValue()
			if err != nil {
				return
			}

			var value Element
			value, err = d.Read()
			if err != nil {
				return
			}

			fields[key] = value
		}

		dataElement = Element{Fields: fields}
	default:
		err = fmt.Errorf("invalid data type")
		return
	}

	return
}

func (d *DataValueReader) readValueType() (valueType ValueType, err error) {
	data, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	valueType = ValueType(data)
	return
}

func (d *DataValueReader) readBool() (value bool, err error) {
	data, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	switch data {
	case 0:
		value = false
	case 1:
		value = true
	}

	return
}

func (d *DataValueReader) readString() (value string, err error) {
	size, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	data := make([]byte, size)
	_, err = d.Reader.Read(data)
	if err != nil {
		return
	}

	value = string(data)
	return
}

func (d *DataValueReader) readBytes(size int) (value []byte, err error) {
	data := make([]byte, size)
	_, err = d.Reader.Read(data)
	if err != nil {
		return
	}

	value = data
	return
}

func (d *DataValueReader) readU8() (value uint8, err error) {
	data, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	value = uint8(data)
	return
}

func (d *DataValueReader) readU16() (value uint16, err error) {
	data, err := d.readBytes(2)
	if err != nil {
		return
	}

	value = binary.BigEndian.Uint16(data)
	return
}

func (d *DataValueReader) readU32() (value uint32, err error) {
	data, err := d.readBytes(4)
	if err != nil {
		return
	}

	value = binary.BigEndian.Uint32(data)
	return
}

func (d *DataValueReader) readU64() (value uint64, err error) {
	data, err := d.readBytes(8)
	if err != nil {
		return
	}

	value = binary.BigEndian.Uint64(data)
	return
}

func (d *DataValueReader) readU128() (value big.Int, err error) {
	data, err := d.readBytes(16)
	if err != nil {
		return
	}

	value.SetBytes(data)
	return
}

func (d *DataValueReader) readHash() (value Hash, err error) {
	data, err := d.readBytes(32)
	if err != nil {
		return
	}

	copy(value[:], data)
	return
}

func (d *DataValueReader) readDataValue() (value Value, err error) {
	valueType, err := d.readValueType()
	if err != nil {
		return
	}

	switch valueType {
	case BoolType:
		value, err = d.readBool()
	case StringType:
		value, err = d.readString()
	case U8Type:
		value, err = d.readU8()
	case U16Type:
		value, err = d.readU16()
	case U32Type:
		value, err = d.readU32()
	case U64Type:
		value, err = d.readU64()
	case U128Type:
		value, err = d.readU128()
	case HashType:
		value, err = d.readHash()
	}

	return
}
