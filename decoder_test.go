package base64

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestDecoder(t *testing.T) {
	for i := 0; i < 1000; i++ {
		srcBytes := make([]byte, i)
		for j := 0; j < i; j++ {
			srcBytes[j] = stdLutSe[j%len(stdLutSe)]
		}
		src := string(srcBytes)

		stdResult, stdErr := base64.StdEncoding.DecodeString(src)
		ownResult, ownErr := StdEncoding.DecodeString(src)
		if (stdErr != nil && ownErr == nil) || (stdErr == nil && ownErr != nil) {
			t.Log("expected:", stdErr)
			t.Log("actual:  ", src)
			t.Fatal()
		} else if stdErr == nil && !bytes.Equal(stdResult, ownResult) {
			t.Log("src:     ", src)
			t.Log("expected:", stdResult)
			t.Log("actual:  ", ownResult)
			t.Fatal()
		}

		stdResult, stdErr = base64.RawStdEncoding.DecodeString(src)
		ownResult, ownErr = RawStdEncoding.DecodeString(src)
		if (stdErr != nil && ownErr == nil) || (stdErr == nil && ownErr != nil) {
			t.Log("expected:", stdErr)
			t.Log("actual:  ", src)
			t.Fatal()
		} else if stdErr == nil && !bytes.Equal(stdResult, ownResult) {
			t.Log("src:     ", src)
			t.Log("expected:", stdResult)
			t.Log("actual:  ", ownResult)
			t.Fatal()
		}
	}

	for i := 1; i < 200; i++ {
		for j := 0; j < 10000; j++ {
			valueBytes := generateRandomBytes(i)

			badValue := b2s(valueBytes)
			_, stdErr := base64.StdEncoding.DecodeString(badValue)
			_, ownErr := StdEncoding.DecodeString(badValue)
			if (stdErr != nil && ownErr == nil) || (stdErr == nil && ownErr != nil) {
				t.Log(stdErr)
				t.Log(valueBytes)
				t.Fatal()
			}

			value := base64.StdEncoding.EncodeToString(valueBytes)
			stdResult, err := base64.StdEncoding.DecodeString(value)
			if err != nil {
				t.Fatal()
			}
			ownResult, err := StdEncoding.DecodeString(value)
			if err != nil {
				t.Fatal()
			}
			if !bytes.Equal(stdResult, ownResult) {
				t.Log("expected:", stdResult)
				t.Log("actual:  ", ownResult)
				t.Fatal()
			}

			value = base64.RawStdEncoding.EncodeToString(valueBytes)
			stdResult, err = base64.RawStdEncoding.DecodeString(value)
			if err != nil {
				t.Fatal()
			}
			ownResult, err = RawStdEncoding.DecodeString(value)
			if err != nil {
				t.Fatal()
			}
			if !bytes.Equal(stdResult, ownResult) {
				t.Log("expected:", stdResult)
				t.Log("actual:  ", ownResult)
				t.Fatal()
			}

			value = base64.URLEncoding.EncodeToString(valueBytes)
			stdResult, err = base64.URLEncoding.DecodeString(value)
			if err != nil {
				t.Fatal()
			}
			ownResult, err = URLEncoding.DecodeString(value)
			if err != nil {
				t.Fatal()
			}
			if !bytes.Equal(stdResult, ownResult) {
				t.Log("expected:", stdResult)
				t.Log("actual:  ", ownResult)
				t.Fatal()
			}

			value = base64.RawURLEncoding.EncodeToString(valueBytes)
			stdResult, err = base64.RawURLEncoding.DecodeString(value)
			if err != nil {
				t.Fatal()
			}
			ownResult, err = RawURLEncoding.DecodeString(value)
			if err != nil {
				t.Fatal()
			}
			if !bytes.Equal(stdResult, ownResult) {
				t.Log("expected:", stdResult)
				t.Log("actual:  ", ownResult)
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

	length = StdEncoding.DecodedLen(len(stdBase64ValueBytes))
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

func BenchmarkDecoder2(b *testing.B) {
	b.ReportAllocs()
	var err error

	length := StdEncoding.DecodedLen(len(stdBase64ValueBytes))
	byteResult = make([]byte, length)
	b.Run("own/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			resultN, err = StdEncoding.Decode(byteResult, stdBase64ValueBytes)
			if err != nil {
				b.Fatal()
			}
		}
	})
}
