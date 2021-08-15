// based on https://github.com/powturbo/Turbo-Base64
package base64

import "errors"

func EncodedLen(n int) int {
	return (n + 2) / 3 * 4
}

func RawEncodedLen(n int) int {
	return (n*8 + 5) / 6
}

func DecodedLen(n int) int {
	return n / 4 * 3
}

func RawDecodedLen(n int) int {
	return n * 6 / 8
}

type Encoding struct {
	lutSe  [64]byte
	lutXe  [4096]uint32
	lutXd0 [256]uint32
	lutXd1 [256]uint32
	lutXd2 [256]uint32
	lutXd3 [256]uint32
	pad    bool
}

func (e *Encoding) Encode(dst []byte, src []byte) {
	var length int
	if e.pad {
		length = EncodedLen(len(src))
	} else {
		length = RawEncodedLen(len(src))
	}
	e.encode(dst, src, uintptr(length))
}

func (e *Encoding) EncodeToBytes(src []byte) []byte {
	var length int
	if e.pad {
		length = EncodedLen(len(src))
	} else {
		length = RawEncodedLen(len(src))
	}
	result := make([]byte, length)
	e.encode(result, src, uintptr(length))
	return result
}

func (e *Encoding) EncodeToString(src []byte) string {
	return b2s(e.EncodeToBytes(src))
}

func (e *Encoding) EncodeString(src string) []byte {
	return e.EncodeToBytes(s2b(src))
}

func (e *Encoding) EncodeStringToString(src string) string {
	return b2s(e.EncodeToBytes(s2b(src)))
}

func (e *Encoding) Decode(dst []byte, src []byte) (int, error) {
	n := e.decode(dst, src)
	if n == 0 {
		return 0, errors.New("wrong base64 data")
	}
	return n, nil
}

func (e *Encoding) DecodeToBytes(src []byte) ([]byte, error) {
	var length int
	if e.pad {
		length = DecodedLen(len(src))
	} else {
		length = RawDecodedLen(len(src))
	}
	result := make([]byte, length)
	n := e.decode(result, src)
	if n == 0 {
		return nil, errors.New("wrong base64 data")
	}
	return result[:n], nil
}

func (e *Encoding) DecodeToString(src []byte) (string, error) {
	result, err := e.DecodeToBytes(src)
	if err != nil {
		return "", err
	}
	return b2s(result), nil
}

func (e *Encoding) DecodeString(src string) ([]byte, error) {
	return e.DecodeToBytes(s2b(src))
}

func (e *Encoding) DecodeStringToString(src string) (string, error) {
	result, err := e.DecodeToBytes(s2b(src))
	if err != nil {
		return "", err
	}
	return b2s(result), nil
}

func NewEncoding(lutSe [64]byte, pad bool) *Encoding {
	lutXe, lutXd0, lutXd1, lutXd2, lutXd3 := makeLuts(lutSe)
	return &Encoding{
		lutSe:  lutSe,
		lutXe:  lutXe,
		lutXd0: lutXd0,
		lutXd1: lutXd1,
		lutXd2: lutXd2,
		lutXd3: lutXd3,
		pad:    pad,
	}
}

var (
	StdEncoding    = NewEncoding(stdLutSe, true)
	RawStdEncoding = NewEncoding(stdLutSe, false)
	URLEncoding    = NewEncoding(urlLutSe, true)
	RawURLEncoding = NewEncoding(urlLutSe, false)
)