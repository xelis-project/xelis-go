package extra_data

import (
	"fmt"
	"math/big"
)

type Element struct {
	Value  Value
	Array  []Element
	Fields map[Value]Element
}

// use this function to convert in a valid map for json.Marshal()
func (d Element) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	if d.Value != nil {
		switch value := d.Value.(type) {
		case big.Int:
			result["value"] = value.String()
		default:
			result["value"] = value
		}
	}

	if d.Array != nil {
		var array []interface{}
		for _, item := range d.Array {
			array = append(array, item.ToMap())
		}

		result["array"] = array
	}

	if d.Fields != nil {
		fields := make(map[string]interface{})
		for key, item := range d.Fields {
			sKey := fmt.Sprintf("%v", key)
			fields[sKey] = item.ToMap()
		}

		result["fields"] = fields
	}

	return result
}
