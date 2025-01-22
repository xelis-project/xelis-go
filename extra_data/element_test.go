package extra_data

import (
	"bytes"
	"testing"
)

func TestElementSerialization(t *testing.T) {
	data := Element{Fields: map[Value]Element{
		"hello": Element{Value: "world"},
	}}

	b, err := data.ToBytes()
	if err != nil {
		t.Fatal(err)
	}

	b2 := []byte{2, 1, 1, 5, 104, 101, 108, 108, 111, 0, 1, 5, 119, 111, 114, 108, 100}
	if !bytes.Equal(b, b2) {
		t.Fail()
	}

	e, err := ElementFromBytes(b)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", e)
}
