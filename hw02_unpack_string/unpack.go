package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	type pair struct {
		r     rune
		count int
	}
	var pairs []pair

	var lastRune rune
	var lastCount int
	var hasLast bool
	lastWasDigit := false

	for _, r := range s {
		if unicode.IsDigit(r) {
			if !hasLast {
				return "", ErrInvalidString
			}
			if lastWasDigit {
				return "", ErrInvalidString
			}
			lastCount = int(r - '0')
			lastWasDigit = true
		} else {
			if hasLast {
				pairs = append(pairs, pair{lastRune, lastCount})
			}
			lastRune = r
			lastCount = 1
			hasLast = true
			lastWasDigit = false
		}
	}
	if hasLast {
		pairs = append(pairs, pair{lastRune, lastCount})
	}

	var sb strings.Builder
	for _, p := range pairs {
		if p.count > 0 {
			sb.WriteString(strings.Repeat(string(p.r), p.count))
		}
	}
	return sb.String(), nil
}
