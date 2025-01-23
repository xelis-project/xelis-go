package data

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"math/big"
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

func TestElementFieldsJSON(t *testing.T) {
	var hash Hash
	shaHash := sha256.Sum256([]byte("asd"))
	copy(hash[:], shaHash[:])

	var bigNumber big.Int
	bigNumber.SetString("2093458230498572039452039485702938475", 10)

	data := Element{
		Fields: map[Value]Element{
			true:                        {Value: false},
			"hello":                     {Value: "world"},
			uint8(21):                   {Value: hash},
			uint16(23452):               {Value: "16"},
			uint32(3567456756):          {Value: "32"},
			uint64(8796789678967899678): {Value: "64"},
			hash:                        {Value: "test_hash"},
			"test_big_number":           {Value: bigNumber},
			"blob":                      {Value: Blob{0, 1, 2, 4, 5}},
			"sub_map": {
				Fields: map[Value]Element{
					"test": {Value: "test"},
				},
			},
		},
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", string(b))
}

func TestElementArrayJSON(t *testing.T) {
	var hash Hash
	shaHash := sha256.Sum256([]byte("asd"))
	copy(hash[:], shaHash[:])

	var bigNumber big.Int
	bigNumber.SetString("2093458230498572039452039485702938475", 10)

	data := Element{
		Array: []Element{
			{Value: true},
			{Value: "test"},
			{Value: uint8(34)},
			{Value: uint16(34523)},
			{Value: uint32(3452305469)},
			{Value: uint64(3452305469567567456)},
			{Value: bigNumber},
			{
				Array: []Element{
					{Value: "sub_array"},
				},
			},
		},
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", string(b))
}

func TestElementValueJSON(t *testing.T) {
	data := Element{Value: 10}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", string(b))
}
