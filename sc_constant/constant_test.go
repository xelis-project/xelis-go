package sc_constant

import (
	"encoding/json"
	"math/big"
	"testing"
)

func TestConstantOptional(t *testing.T) {
	data, err := json.MarshalIndent(Optional(DefaultString("testing")), "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}

func TestConstantArray(t *testing.T) {
	u128 := new(big.Int)
	u128.SetString("340282366920938463463374607431768211455", 10)

	u256 := new(big.Int)
	u256.SetString("115792089237316195423570985008687907853269984665640564039457584007913129659255", 10)

	data, err := json.MarshalIndent(Array([]Constant{
		DefaultU8(100),
		DefaultU16(10023),
		DefaultU32(143255),
		DefaultU64(23452345),
		DefaultU128(u128),
		DefaultU256(u256),
		DefaultBool(false),
		DefaultString("asdasd"),
	}), "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}

func TestValueRange(t *testing.T) {
	data, err := json.MarshalIndent(ValueRange(
		ValueU16(100),
		ValueU16(200),
	), "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
