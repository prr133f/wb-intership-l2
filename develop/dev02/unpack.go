package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
)

func UnpackString(s string) (string, error) {
	var sb strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			if i == 0 || unicode.IsDigit(runes[i-1]) {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(runes[i]))
			sb.WriteString(strings.Repeat(string(runes[i-1]), count-1))
		} else {
			sb.WriteRune(runes[i])
		}
	}

	return sb.String(), nil
}
