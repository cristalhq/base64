package base64

import (
	"bytes"
	stdbase64 "encoding/base64"
	"fmt"

	"github.com/cristalhq/base64"
)

func Fuzz(data []byte) int {
	stdStringResult := stdbase64.StdEncoding.EncodeToString(data)
	ownStringResult := base64.StdEncoding.EncodeToString(data)
	if stdStringResult != ownStringResult {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdStringResult)
		fmt.Println("actual:  ", ownStringResult)
		panic("mismatch")
	}
	stdStringResult = stdbase64.RawStdEncoding.EncodeToString(data)
	ownStringResult = base64.RawStdEncoding.EncodeToString(data)
	if stdStringResult != ownStringResult {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdStringResult)
		fmt.Println("actual:  ", ownStringResult)
		panic("mismatch")
	}
	stdStringResult = stdbase64.URLEncoding.EncodeToString(data)
	ownStringResult = base64.URLEncoding.EncodeToString(data)
	if stdStringResult != ownStringResult {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdStringResult)
		fmt.Println("actual:  ", ownStringResult)
		panic("mismatch")
	}
	stdStringResult = stdbase64.RawURLEncoding.EncodeToString(data)
	ownStringResult = base64.RawURLEncoding.EncodeToString(data)
	if stdStringResult != ownStringResult {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdStringResult)
		fmt.Println("actual:  ", ownStringResult)
		panic("mismatch")
	}

	if bytes.ContainsRune(data, '\r') || bytes.ContainsRune(data, '\n') {
		return -1
	}

	score := 0
	stringData := string(data)
	stdResult, stdErr := stdbase64.StdEncoding.DecodeString(stringData)
	if stdErr == nil {
		score = 1
	}
	ownResult, ownErr := base64.StdEncoding.DecodeString(stringData)
	if (stdErr != nil && ownErr == nil) || (stdErr == nil && ownErr != nil) {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdErr)
		fmt.Println("actual:  ", ownErr)
		panic("mismatch")
	} else if stdErr == nil && !bytes.Equal(stdResult, ownResult) {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdResult)
		fmt.Println("actual:  ", ownResult)
		panic("mismatch")
	}
	stdResult, stdErr = stdbase64.RawStdEncoding.DecodeString(stringData)
	if stdErr == nil {
		score = 1
	}
	ownResult, ownErr = base64.RawStdEncoding.DecodeString(stringData)
	if (stdErr != nil && ownErr == nil) || (stdErr == nil && ownErr != nil) {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdErr)
		fmt.Println("actual:  ", ownErr)
		panic("mismatch")
	} else if stdErr == nil && !bytes.Equal(stdResult, ownResult) {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdResult)
		fmt.Println("actual:  ", ownResult)
		panic("mismatch")
	}
	stdResult, stdErr = stdbase64.URLEncoding.DecodeString(stringData)
	if stdErr == nil {
		score = 1
	}
	ownResult, ownErr = base64.URLEncoding.DecodeString(stringData)
	if (stdErr != nil && ownErr == nil) || (stdErr == nil && ownErr != nil) {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdErr)
		fmt.Println("actual:  ", ownErr)
		panic("mismatch")
	} else if stdErr == nil && !bytes.Equal(stdResult, ownResult) {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdResult)
		fmt.Println("actual:  ", ownResult)
		panic("mismatch")
	}
	stdResult, stdErr = stdbase64.RawURLEncoding.DecodeString(stringData)
	if stdErr == nil {
		score = 1
	}
	ownResult, ownErr = base64.RawURLEncoding.DecodeString(stringData)
	if (stdErr != nil && ownErr == nil) || (stdErr == nil && ownErr != nil) {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdErr)
		fmt.Println("actual:  ", ownErr)
		panic("mismatch")
	} else if stdErr == nil && !bytes.Equal(stdResult, ownResult) {
		fmt.Println("value:   ", data)
		fmt.Println("expected:", stdResult)
		fmt.Println("actual:  ", ownResult)
		panic("mismatch")
	}
	return score
}
