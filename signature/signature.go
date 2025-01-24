package signature

import (
	"encoding/hex"

	"github.com/gtank/ristretto255"
	"golang.org/x/crypto/sha3"
)

// based on https://github.com/xelis-project/xelis-blockchain/blob/master/xelis_common/src/crypto/elgamal/signature.rs#L29

var ristretto_basepoint_compressed = []byte{
	0xe2, 0xf2, 0xae, 0x0a, 0x6a, 0xbc, 0x4e, 0x71, 0xa8, 0x84, 0xa9, 0x61, 0xc5, 0x00, 0x51, 0x5f,
	0x58, 0xe3, 0x0b, 0x6a, 0xa5, 0x82, 0xdd, 0x8d, 0xb6, 0xa6, 0x59, 0x45, 0xe0, 0x8d, 0x2d, 0x76,
}

func Verify2(publicKey string, signature string, data []byte) (valid bool, err error) {
	slice, err := hex.DecodeString(publicKey)
	if err != nil {
		return
	}

	var pk [32]byte
	copy(pk[:], slice)

	return Verify(pk, signature, data)
}

func Verify(publicKey [32]byte, signature string, data []byte) (valid bool, err error) {
	bBlindingBytes := createBlinding()
	sBytes, eBytes, err := splitSignature(signature)
	if err != nil {
		return
	}

	b := &ristretto255.Element{}
	b.FromUniformBytes(bBlindingBytes)

	k := &ristretto255.Element{}
	k.Decode(publicKey[:])

	s := &ristretto255.Scalar{}
	s.Decode(sBytes)

	e := &ristretto255.Scalar{}
	e.Decode(eBytes)

	// b * s + k * -e
	e.Negate(e)
	b.ScalarMult(s, b)
	k.ScalarMult(e, k)
	r := b.Add(k, b)

	point, err := hashAndPointToScalar(publicKey, data, r)
	if err != nil {
		return
	}

	e.Negate(e)
	if e.Equal(point) == 1 {
		valid = true
	}

	return
}

func splitSignature(signature string) (s []byte, e []byte, err error) {
	data := make([]byte, 64)
	_, err = hex.Decode(data, []byte(signature))
	s = data[:32]
	e = data[32:]
	return
}

func createBlinding() []byte {
	hash := sha3.New512()
	hash.Write(ristretto_basepoint_compressed)
	return hash.Sum(nil)
}

func hashAndPointToScalar(publicKey [32]byte, data []byte, point *ristretto255.Element) (scalar *ristretto255.Scalar, err error) {
	hash := sha3.New512()
	hash.Write(publicKey[:])
	hash.Write(data)
	hash.Write(point.Encode([]byte{}))
	hashed := hash.Sum(nil)

	scalar = &ristretto255.Scalar{}
	scalar.FromUniformBytes(hashed)
	return
}
