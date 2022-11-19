package perf

import (
	"strings"
)

func Contains(s1, s2 string) bool {
	lenS1 := len(s1)
	lenS2 := len(s2)
	for i := 0; i < lenS1-lenS2+1; i++ {
		subS1 := s1[i : i+lenS2]
		if subS1 == s2 {
			return true
		}
	}
	return false
}

func isPalindromeV1(x string) bool {
	if x == "" {
		return true
	}
	lenX := len(x)
	for i, j := lenX-1, 0; j < i; i, j = i-1, j+1 {
		if x[i] != x[j] {
			return false
		}
	}
	return true
}

func isPalindromeV2(x string) bool {
	if x == "" {
		return true
	}

	lenX := len(x)
	for i := 0; i < lenX; i++ {
		if x[i] != x[lenX-i-1] {
			return false
		}
	}
	return true
}

// Typically suggested code with runes
func ReverseWithRune(str string) string {
	runes := []rune(str)
	for i, j := len(runes)-1, 0; j < i; i, j = i-1, j+1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// uses strings.Builder for no-copy
func ReverseWithBuilder(str string) string {
	var ret strings.Builder
	strLen := len(str)
	ret.Grow(strLen)
	for i := strLen - 1; i >= 0; i-- {
		ret.WriteByte(str[i])
	}
	return ret.String()
}

// Plain byte array and tranverse string in reverse
func ReverseWithByteArray(str string) string {
	strLen := len(str)
	ret := make([]byte, strLen)
	for i := strLen - 1; i >= 0; i-- {
		ret[strLen-i-1] = str[i]
	}
	return string(ret)
}
