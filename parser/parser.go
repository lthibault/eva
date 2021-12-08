package parser

import (
	"io"
	"strings"
)

// Parse the program.
//
// Program
//  : NumericLiteral
//  ;
//
func Parse(r io.RuneScanner) (NumericLiteral, error) {
	return ParseNumericLiteral(r)
}

func ParseNumericLiteral(r io.RuneScanner) (NumericLiteral, error) {
	var (
		s    strings.Builder
		char rune
		err  error
	)

	for err == nil {
		if char, _, err = r.ReadRune(); err == nil {
			s.WriteRune(char)
		}
	}

	return NumericLiteral{s.String()}, err
}

type NumericLiteral struct {
	Value string
}

func (NumericLiteral) Type() string { return "NumericLiteral" }
