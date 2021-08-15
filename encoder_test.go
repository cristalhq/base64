package base64

import (
	"encoding/base64"
	"testing"
)

func TestEncoder(t *testing.T) {
	for i := 1; i < 200; i++ {
		for j := 0; j < 10000; j++ {
			bytes := generateRandomBytes(i)

			stdResult := base64.StdEncoding.EncodeToString(bytes)
			ownResult := StdEncoding.EncodeToString(bytes)
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
