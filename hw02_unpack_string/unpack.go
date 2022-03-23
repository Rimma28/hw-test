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
	var test string
	var count int
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			if previoseLetter == "" {
				err := ErrInvalidString
				return result, err
			}
			count, _ = strconv.Atoi(string(ch))
			if count > 0 {
				test = strings.Repeat(previoseLetter, count)
				result = result + test
			}
			previoseLetter = ""
		} else {
			if previoseLetter != "" {
				result = result + previoseLetter
			}
			previoseLetter = string(ch)
		}
	}
	result = result + previoseLetter
	return result, nil
}
