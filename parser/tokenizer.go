package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Kind int

const (
	Int Kind = iota
	String
)

func (k Kind) String() string {
	switch k {
	case Int:
		return "int"

	case String:
		return "string"

	default:
		panic(fmt.Errorf("unrecognized kind '%d'", k))
	}
}

func ReadToken(r io.RuneReader) (Token, error) {
	init, _, err := r.ReadRune()
	if err != nil {
		return Token{}, err
	}

	if unicode.IsNumber(init) {
		return readInt(r, init)
	}

	if init == '"' || init == '\'' {
		return readString(r, init)
	}

	return Token{}, fmt.Errorf("syntax error: '%c'", init)
}

func readInt(r io.RuneReader, init rune) (Token, error) {
	var (
		err error
		b   strings.Builder
	)

	for err == nil {
		b.WriteRune(init)

		init, _, err = r.ReadRune()
		if (err != nil && !errors.Is(err, io.EOF)) || !unicode.IsNumber(init) {
			break
		}
	}

	// will return EOF on next call
	if errors.Is(err, io.EOF) && b.Len() > 0 {
		err = nil
	}

	return Token{
		Kind:  Int,
		Value: b.String(),
	}, err
}

func readString(r io.RuneReader, quote rune) (Token, error) {
	var (
		err  error
		b    strings.Builder
		char rune
	)

	for b.WriteRune(quote); err == nil; b.WriteRune(char) {
		if char, _, err = r.ReadRune(); err == nil && char == quote {
			b.WriteRune(char)
			break
		}
	}

	return Token{
		Kind:  String,
		Value: b.String(),
	}, err
}

type Token struct {
	Kind  Kind
	Value string
}
