package extra_data

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"reflect"
	"testing"
)

func writeDataElement(dataElement Element) (data []byte, err error) {
	var buf bytes.Buffer
	dataValueWriter := ValueWriter{Writer: &buf}
	err = dataValueWriter.Write(dataElement)
	if err != nil {
		return
	}

	data = buf.Bytes()
	return
}

func readDataElement(data []byte) (dataElement Element, err error) {
	dataValueReader := ValueReader{Reader: bytes.NewReader(data)}
	dataElement, err = dataValueReader.Read()
	return
}

func TestDataElementBoolValue(t *testing.T) {
	dataElement := Element{Value: true}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %t, got %t", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementStringValue(t *testing.T) {
	dataElement := Element{Value: "test"}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %s, got %s", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU8Value(t *testing.T) {
	dataElement := Element{Value: uint8(122)}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU16Value(t *testing.T) {
	dataElement := Element{Value: uint16(34566)}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU32Value(t *testing.T) {
	dataElement := Element{Value: uint32(6767456)}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU64Value(t *testing.T) {
	dataElement := Element{Value: uint64(345765875678)}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU128Value(t *testing.T) {
	var nbr big.Int
	nbr.SetString("35467456745674956794567", 10)

	dataElement := Element{Value: nbr}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	v1 := dataElement.Value.(big.Int)
	v2 := dataElement.Value.(big.Int)

	if v1.Cmp(&v2) != 0 {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementHashValue(t *testing.T) {
	var hash Hash

	shaHash := sha256.Sum256([]byte("asd"))
	copy(hash[:], shaHash[:])

	dataElement := Element{Value: hash}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	v1 := dataElement.Value.(Hash)
	v2 := dataElementCopy.Value.(Hash)

	if !bytes.Equal(v1[:], v2[:]) {
		t.Fail()
		t.Logf("Expected %s, got %s", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementBlobValue(t *testing.T) {
	var arr Blob
	arr = append(arr, 0, 1, 2, 3, 4, 5)

	dataElement := Element{Value: arr}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	v1 := dataElement.Value.(Blob)
	v2 := dataElementCopy.Value.(Blob)

	if !reflect.DeepEqual(v1, v2) {
		t.Fail()
		t.Logf("Expected %s, got %s", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementArray(t *testing.T) {
	array := []Element{
		{Value: true},
		{Value: "test"},
	}

	dataElement := Element{Array: array}

	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	for i, item := range array {
		if dataElementCopy.Array[i].Value != item.Value {
			t.Fail()
			t.Logf("Expected %v, got %v", item.Value, dataElementCopy.Array[i].Value)
		}
	}
}

func TestDataElementFields(t *testing.T) {
	fields := make(map[Value]Element, 0)
	fields["hello"] = Element{Value: "world"}

	dataElement := Element{Fields: fields}

	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	for key, element := range fields {
		if dataElementCopy.Fields[key].Value != element.Value {
			t.Fail()
			t.Logf("Expected %v, got %v", element.Value, dataElementCopy.Fields[key].Value)
		}
	}
}

func TestLongStringMaxLimit(t *testing.T) {
	// max 255 bytes for string
	_, err := writeDataElement(Element{Value: "woenrbowirentboiejwrntbpoijewnrtbpenrptbjnepritjbnperijtnbpijewnrtbpjnerptbnjperkjtbnperkjtnbpsdfgsergwngio453gn45oign345iogjnwosiwejrngwpo34i5ny3[45oyhi3n4p5[hokn3p4o5nhekjrntbpkjewnrtpbkjnwerptkbjnpwkrjntbperkjntbpkwerntpbjkenrptbkjnwpekjnrwkpenrfpbknweprkbjnwperkbjn"})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
