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
		err error
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde", err: nil},
		{input: "abccd", expected: "abccd", err: nil},
		{input: "", expected: "", err: nil},
		{input: "aaa0b", expected: "aab", err: nil},
		{input: "🙃0", expected: "", err: nil},
		{input: "aaф0b", expected: "aab", err: nil},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
		// дополнительные успешные тесты
        {input: "a0", expected: "", err: nil},
        {input: "a0b", expected: "b", err: nil},
        {input: "a1b2c3", expected: "abbccc", err: nil},
        {input: "я9", expected: "яяяяяяяяя", err: nil},
        {input: "😀3", expected: "😀😀😀", err: nil},
        {input: "a0b0c", expected: "c", err: nil},
        {input: "a5", expected: "aaaaa", err: nil},
		// дополнительные ошибочные тесты
		{input: "3abc", expected: "", err: ErrInvalidString},
        {input: "45", expected: "", err: ErrInvalidString},
        {input: "aaa10b", expected: "", err: ErrInvalidString},
        {input: "a00", expected: "", err: ErrInvalidString},
        {input: "0", expected: "", err: ErrInvalidString},
        {input: "1a", expected: "", err: ErrInvalidString},
	}

	for _, tt := range tests {
        result, err := Unpack(tt.input)
        if tt.err != nil {
            require.ErrorIs(t, err, tt.err)
        } else {
            require.NoError(t, err)
            require.Equal(t, tt.expected, result)
        }
    }
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
