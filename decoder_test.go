package base64

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoder(t *testing.T) {
	for i := 1; i < 200; i++ {
		for j := 0; j < 10000; j++ {
			bytes := generateRandomBytes(i)

			value := base64.StdEncoding.EncodeToString(bytes)
			stdResult, err := base64.StdEncoding.DecodeString(value)
			if !assert.NoError(t, err) {
				t.Fatal()
			}
			ownResult, err := StdEncoding.DecodeString(value)
			if !assert.NoError(t, err) {
				t.Fatal()
			}
			if !assert.Equal(t, stdResult, ownResult) {
				t.Fatal()
			}

			value = base64.RawStdEncoding.EncodeToString(bytes)
			stdResult, err = base64.RawStdEncoding.DecodeString(value)
			if !assert.NoError(t, err) {
				t.Fatal()
			}
			ownResult, err = RawStdEncoding.DecodeString(value)
			if !assert.NoError(t, err) {
				t.Fatal()
			}
			if !assert.Equal(t, stdResult, ownResult) {
				t.Fatal()
			}

			value = base64.URLEncoding.EncodeToString(bytes)
			stdResult, err = base64.URLEncoding.DecodeString(value)
			if !assert.NoError(t, err) {
				t.Fatal()
			}
			ownResult, err = URLEncoding.DecodeString(value)
			if !assert.NoError(t, err) {
				t.Fatal()
			}
			if !assert.Equal(t, stdResult, ownResult) {
				t.Fatal()
			}

			value = base64.RawURLEncoding.EncodeToString(bytes)
			stdResult, err = base64.RawURLEncoding.DecodeString(value)
			if !assert.NoError(t, err) {
				t.Fatal()
			}
			ownResult, err = RawURLEncoding.DecodeString(value)
			if !assert.NoError(t, err) {
				t.Fatal()
			}
			if !assert.Equal(t, stdResult, ownResult) {
				t.Fatal()
			}
		}
	}
}

func BenchmarkDecoder(b *testing.B) {
	b.ReportAllocs()
	var err error

	length := base64.StdEncoding.DecodedLen(len(stdBase64ValueBytes))
	byteResult = make([]byte, length)
	b.Run("std/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			resultN, err = base64.StdEncoding.Decode(byteResult, stdBase64ValueBytes)
			if err != nil {
				b.Fatal()
			}
		}
	})
	b.Run("std/DecodeString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			byteResult, err = base64.StdEncoding.DecodeString(stdBase64ValueString)
			if err != nil {
				b.Fatal()
			}
		}
	})

	length = DecodedLen(len(stdBase64ValueBytes))
	byteResult = make([]byte, length)
	b.Run("own/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			resultN, err = StdEncoding.Decode(byteResult, stdBase64ValueBytes)
			if err != nil {
				b.Fatal()
			}
		}
	})
	b.Run("own/DecodeString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			byteResult, err = StdEncoding.DecodeString(stdBase64ValueString)
			if err != nil {
				b.Fatal()
			}
		}
	})
}
