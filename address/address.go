package address

import (
	"bytes"
	"errors"
	"fmt"

	d "github.com/xelis-project/xelis-go-sdk/data"
)

var PrefixAddress string = "xel"
var TestnetPrefixAddress string = "xet"

var ExtraDataLimit = 1024

var ErrIntegratedDataLimit = errors.New("invalid data in integrated address, maximum size reached")

type Address struct {
	publicKey    []byte
	isMainnet    bool
	isIntegrated bool
	extraData    *d.Element
}

func NewAddressFromData(data []byte, hrp string) (addr *Address, err error) {
	reader := bytes.NewReader(data)

	publicKey := make([]byte, 32)
	_, err = reader.Read(publicKey)
	if err != nil {
		return
	}

	addrType, err := reader.ReadByte()
	if err != nil {
		return
	}

	integrated := false
	var extraData d.Element

	switch addrType {
	case 0:
		// do nothing
	case 1:
		integrated = true

		dataValueReader := &d.ValueReader{Reader: reader}
		extraData, err = dataValueReader.Read()
		if err != nil {
			return
		}

		if reader.Size() > int64(ExtraDataLimit) {
			err = ErrIntegratedDataLimit
			return
		}
	default:
		err = fmt.Errorf("invalid address type")
		return
	}

	addr = &Address{
		isMainnet:    hrp == PrefixAddress,
		publicKey:    publicKey,
		isIntegrated: integrated,
		extraData:    &extraData,
	}

	return
}

func NewAddressFromString(address string) (addr *Address, err error) {
	hrp, decoded, err := decode(address)
	if err != nil {
		return
	}

	if hrp != PrefixAddress && hrp != TestnetPrefixAddress {
		return
	}

	bits, err := convertBits(decoded, 5, 8, false)
	if err != nil {
		return
	}

	addr, err = NewAddressFromData(bits, hrp)
	if err != nil {
		return
	}

	return
}

func IsValidAddress(address string) (valid bool, err error) {
	_, err = NewAddressFromString(address)
	if err == nil {
		valid = true
	}

	return
}

func (a *Address) IsMainnet() bool {
	return a.isMainnet
}

func (a *Address) IsIntegrated() bool {
	return a.isIntegrated
}

func (a *Address) GetPublicKey() []byte {
	return a.publicKey
}

func (a *Address) GetExtraData() *d.Element {
	return a.extraData
}

func (a *Address) SetExtraData(data *d.Element) {
	if data != nil {
		a.isIntegrated = true
		a.extraData = data
	} else {
		a.isIntegrated = false
		a.extraData = nil
	}
}

// Same as using SetExtraData(nil)
func (a *Address) ClearExtraData() {
	a.isIntegrated = false
	a.extraData = nil
}

func (a *Address) Format() (addr string, err error) {
	var buf bytes.Buffer
	_, err = buf.Write(a.publicKey)
	if err != nil {
		return
	}

	if a.isIntegrated {
		_, err = buf.Write([]byte{1})
		if err != nil {
			return
		}

		var extraData []byte
		extraData, err = a.extraData.ToBytes()
		if err != nil {
			return
		}

		_, err = buf.Write(extraData)
		if err != nil {
			return
		}
	} else {
		_, err = buf.Write([]byte{0})
		if err != nil {
			return
		}
	}

	bits, err := convertBits(buf.Bytes(), 8, 5, true)
	if err != nil {
		return
	}

	hrp := PrefixAddress
	if !a.isMainnet {
		hrp = TestnetPrefixAddress
	}

	addr, err = encode(hrp, bits)
	return
}
