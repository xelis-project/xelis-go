package extra_data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
)

type ElementType int

var ElementValueType ElementType = 0
var ElementArrayType ElementType = 1
var ElementFieldsType ElementType = 2

type Element struct {
	Value  Value
	Array  []Element
	Fields map[Value]Element
}

func ElementFromBytes(data []byte) (dataElement Element, err error) {
	reader := ValueReader{Reader: bytes.NewReader(data)}
	dataElement, err = reader.Read()
	return
}

func (d Element) ToBytes() (data []byte, err error) {
	var buf bytes.Buffer
	writer := ValueWriter{Writer: &buf}
	err = writer.Write(d)
	if err != nil {
		return
	}

	data = buf.Bytes()
	return
}

func (d Element) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ToMap())
}

// use this function to convert to a valid interface for json.Marshal()
func (d Element) ToMap() interface{} {
	if d.Value != nil {
		switch value := d.Value.(type) {
		case big.Int:
			return value.String()
		default:
			return value
		}
	} else if d.Array != nil {
		var array []interface{}
		for _, item := range d.Array {
			array = append(array, item.ToMap())
		}

		return array
	} else if d.Fields != nil {
		fields := make(map[string]interface{})
		for key, item := range d.Fields {
			sKey := fmt.Sprintf("%v", key)
			fields[sKey] = item.ToMap()
		}

		return fields
	}

	return nil
}
