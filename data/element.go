package data

import (
	"bytes"
	"encoding/json"
	"errors"
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
	m, err := d.ToMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(m)
}

func (d *Element) validate() (eType ElementType, err error) {
	count := 0

	if d.Value != nil {
		count++
		eType = ElementValueType
	}

	if d.Array != nil {
		count++
		eType = ElementArrayType
	}

	if d.Fields != nil {
		count++
		eType = ElementFieldsType
	}

	if count > 1 {
		err = errors.New("only one field (Value, Array, or Fields) must be set")
		return
	}

	return
}

// use this function to convert to a valid interface for json.Marshal()
func (d Element) ToMap() (data interface{}, err error) {
	field, err := d.validate()
	if err != nil {
		return
	}

	switch field {
	case ElementValueType:
		switch value := d.Value.(type) {
		case big.Int:
			data = value.String()
		default:
			data = value
		}
	case ElementArrayType:
		var array []interface{}
		for _, item := range d.Array {
			var m interface{}
			m, err = item.ToMap()
			if err != nil {
				return
			}

			array = append(array, m)
		}

		data = array
	case ElementFieldsType:
		fields := make(map[string]interface{})
		for key, item := range d.Fields {
			sKey := fmt.Sprintf("%v", key)
			var m interface{}
			m, err = item.ToMap()
			if err != nil {
				return
			}

			fields[sKey] = m
		}

		data = fields
	}

	return
}
