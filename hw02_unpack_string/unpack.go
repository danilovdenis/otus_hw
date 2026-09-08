package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const backslashCode = 92

func Unpack(str string) (string, error) {
	var builder strings.Builder

	runes := []rune(str)

	for i := 0; i < len(runes); {
		var ch rune

		if backslashCode == runes[i] {
			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}

			next := runes[i+1]

			if !unicode.IsDigit(next) && backslashCode != next {
				return "", ErrInvalidString
			}

			ch = next
			i += 2
		} else if unicode.IsDigit(runes[i]) {
			return "", ErrInvalidString
		} else {
			ch = runes[i]
			i++
		}

		cnt := 1
		if i < len(runes) && unicode.IsDigit(runes[i]) {
			cnt, _ = strconv.Atoi(string(runes[i]))
			i++
		}

		builder.WriteString(strings.Repeat(string(ch), cnt))
	}

	return builder.String(), nil
}
