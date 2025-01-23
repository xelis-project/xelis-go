package data

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"
)

func ErrUnsupportedValue(value Value) error {
	return fmt.Errorf("unsupported value type %v", value)
}

var MaxStringSize = 255
var ErrMaxStringSize = errors.New("string max limit is 255 bytes")
var MaxBlobSize = 65535
var ErrMaxBlobSize = errors.New("blob max size is 65535 bytes")

type ValueWriter struct {
	Writer io.Writer
}

func (d *ValueWriter) Write(dataElement Element) (err error) {
	if dataElement.Value != nil {
		err = d.writeByte(byte(ElementValueType))
		if err != nil {
			return
		}

		err = d.writeValue(dataElement.Value)
		if err != nil {
			return
		}
	}

	if dataElement.Array != nil {
		err = d.writeByte(byte(ElementArrayType))
		if err != nil {
			return
		}

		err = d.writeByte(byte(len(dataElement.Array)))
		if err != nil {
			return
		}

		for _, item := range dataElement.Array {
			err = d.Write(item)
			if err != nil {
				return
			}
		}
	}

	if dataElement.Fields != nil {
		err = d.writeByte(byte(ElementFieldsType))
		if err != nil {
			return
		}

		err = d.writeByte(byte(len(dataElement.Fields)))
		if err != nil {
			return
		}

		for key, item := range dataElement.Fields {
			err = d.writeValue(key)
			if err != nil {
				return
			}

			err = d.Write(item)
			if err != nil {
				return
			}
		}
	}

	return
}

func (d *ValueWriter) writeByte(value byte) (err error) {
	data := make([]byte, 1)
	data[0] = value

	_, err = d.Writer.Write(data)
	if err != nil {
		return
	}

	return
}

func (d *ValueWriter) writeBlob(value Blob) (err error) {
	size := uint32(len(value))
	if size > 65535 {
		err = ErrMaxBlobSize
		return
	}

	err = binary.Write(d.Writer, binary.BigEndian, size)
	if err != nil {
		return
	}

	// blob is an array of bytes so simple write it :)
	_, err = d.Writer.Write(value)
	if err != nil {
		return
	}

	return
}

func (d *ValueWriter) writeBool(value bool) (err error) {
	data := make([]byte, 1)

	if value {
		data[0] = 1
	} else {
		data[0] = 0
	}

	_, err = d.Writer.Write(data)
	if err != nil {
		return
	}

	return
}

func (d *ValueWriter) writeString(value string) (err error) {
	buf := bytes.NewBufferString(value)

	if buf.Len() > MaxStringSize {
		err = ErrMaxStringSize
		return
	}

	size := byte(buf.Len())
	err = d.writeByte(size)
	if err != nil {
		return
	}

	_, err = d.Writer.Write(buf.Bytes())
	if err != nil {
		return
	}

	return
}

func (d *ValueWriter) writeU16(value uint16) (err error) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, value)
	_, err = d.Writer.Write(data)
	return
}

func (d *ValueWriter) writeU32(value uint32) (err error) {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, value)
	_, err = d.Writer.Write(data)
	return
}

func (d *ValueWriter) writeU64(value uint64) (err error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, value)
	_, err = d.Writer.Write(data)
	return
}

func (d *ValueWriter) writeU128(value big.Int) (err error) {
	_, err = d.Writer.Write(value.Bytes())
	return
}

func (d *ValueWriter) writeValue(value Value) (err error) {
	switch value := value.(type) {
	case bool:
		err = d.writeByte(byte(BoolType))
		if err != nil {
			return
		}

		err = d.writeBool(value)
	case string:
		err = d.writeByte(byte(StringType))
		if err != nil {
			return
		}

		err = d.writeString(value)
	case uint8:
		err = d.writeByte(byte(U8Type))
		if err != nil {
			return
		}

		err = d.writeByte(byte(value))
	case uint16:
		err = d.writeByte(byte(U16Type))
		if err != nil {
			return
		}

		err = d.writeU16(value)
	case uint32:
		err = d.writeByte(byte(U32Type))
		if err != nil {
			return
		}

		err = d.writeU32(value)
	case uint64:
		err = d.writeByte(byte(U64Type))
		if err != nil {
			return
		}

		err = d.writeU64(value)
	case big.Int:
		err = d.writeByte(byte(U128Type))
		if err != nil {
			return
		}

		err = d.writeU128(value)
	case Hash:
		err = d.writeByte(byte(HashType))
		if err != nil {
			return
		}

		_, err = d.Writer.Write(value[:])
	case Blob:
		err = d.writeByte(byte(BlobType))
		if err != nil {
			return
		}

		err = d.writeBlob(value)
	default:
		err = ErrUnsupportedValue(value)
		return
	}

	return
}
