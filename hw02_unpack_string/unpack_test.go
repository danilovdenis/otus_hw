package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		{input: "😀3", expected: "😀😀😀"},
		{input: "日本2語", expected: "日本本語"},
		{input: "ä2ö3", expected: "ääööö"},
		{input: "я2а3", expected: "яяааа"},
		{input: "キ2ャ3", expected: "キキャャャ"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "a\t2b", expected: "a\t\tb"},
		{input: "a\r2b", expected: "a\r\rb"},
		{input: "ф2я3", expected: "ффяяя"},
		{input: "🙂2😀3", expected: "🙂🙂😀😀😀"},
		{input: "тест2!", expected: "тестт!"},
		// доп задание
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{
		"3abc",
		"45",
		"aaa10b",
		"3привет",
		`qw\e`,
		`qw\n`,
		`qw\`,
		`qw\\\`,
	}
	for _, tc := range invalidStrings {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
