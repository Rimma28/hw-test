package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result string
	var previoseLetter string
	var count int
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			if previoseLetter == "" {
				err := ErrInvalidString
				return result, err
			}
			count, _ = strconv.Atoi(string(ch))
			if count > 0 {
				result += strings.Repeat(previoseLetter, count)
			}
			previoseLetter = ""
		} else {
			if previoseLetter != "" {
				result += previoseLetter
			}
			previoseLetter = string(ch)
		}
	}
	result += previoseLetter
	return result, nil
}
