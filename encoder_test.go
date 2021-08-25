package base64

import (
	"encoding/base64"
	"testing"
	"unsafe"
)

func TestEncoder(t *testing.T) {
	stdResult := base64.StdEncoding.EncodeToString(nil)
	ownResult := StdEncoding.EncodeToString(nil)
	if stdResult != ownResult {
		t.Log("expected:", stdResult)
		t.Log("actual:  ", ownResult)
		t.Fatal()
	}

	stdResult = base64.StdEncoding.EncodeToString([]byte{})
	ownResult = StdEncoding.EncodeToString([]byte{})
	if stdResult != ownResult {
		t.Log("expected:", stdResult)
		t.Log("actual:  ", ownResult)
		t.Fatal()
	}

	for i := 1; i < 200; i++ {
		for j := 0; j < 10000; j++ {
			bytes := generateRandomBytes(i)

			for _, encoding := range []*Encoding{StdEncoding, RawStdEncoding, URLEncoding, RawURLEncoding} {
				length := encoding.EncodedLen(len(bytes))
				if length == 0 {
					continue
				}
				result := make([]byte, length, length+10)
				encoding.encode(result, bytes, uintptr(length))
				(*sliceHeader)(unsafe.Pointer(&result)).len = length + 10
				for _, b := range result[length:] {
					if b != 0 {
						t.Fatal("out of bounds")
					}
				}
			}

			stdResult = base64.StdEncoding.EncodeToString(bytes)
			ownResult = StdEncoding.EncodeToString(bytes)
			if stdResult != ownResult {
				t.Log("expected:", stdResult)
				t.Log("actual:  ", ownResult)
				t.Fatal()
			}

			stdResult = base64.RawStdEncoding.EncodeToString(bytes)
			ownResult = RawStdEncoding.EncodeToString(bytes)
			if stdResult != ownResult {
				t.Log("expected:", stdResult)
				t.Log("actual:  ", ownResult)
				t.Fatal()
			}

			stdResult = base64.URLEncoding.EncodeToString(bytes)
			ownResult = URLEncoding.EncodeToString(bytes)
			if stdResult != ownResult {
				t.Log("expected:", stdResult)
				t.Log("actual:  ", ownResult)
				t.Fatal()
			}

			stdResult = base64.RawURLEncoding.EncodeToString(bytes)
			ownResult = RawURLEncoding.EncodeToString(bytes)
			if stdResult != ownResult {
				t.Log("expected:", stdResult)
				t.Log("actual:  ", ownResult)
				t.Fatal()
			}
		}
	}
}

func BenchmarkEncoder(b *testing.B) {
	b.ReportAllocs()

	length := base64.StdEncoding.EncodedLen(len(valueBytes))
	byteResult = make([]byte, length)
	b.Run("std/Encode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			base64.StdEncoding.Encode(byteResult, valueBytes)
		}
	})
	b.Run("std/EncodeToString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			stringResult = base64.StdEncoding.EncodeToString(valueBytes)
		}
	})

	length = StdEncoding.EncodedLen(len(valueBytes))
	byteResult = make([]byte, length)
	b.Run("own/Encode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			StdEncoding.Encode(byteResult, valueBytes)
		}
	})
	b.Run("own/EncodeToString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			stringResult = StdEncoding.EncodeToString(valueBytes)
		}
	})
}
