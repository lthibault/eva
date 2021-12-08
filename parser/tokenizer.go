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

func ReadToken(r io.RuneScanner) (Token, error) {
	init, err := discardWhitespace(r)
	if err != nil {
		return Token{}, err
	}

	if unicode.IsNumber(init) {
		return readInt(r, init)
	}

	if init == '"' || init == '\'' {
		return readString(r, init)
	}

	// comment?
	if init == '/' {
		next, _, err := r.ReadRune()
		if err != nil {
			return Token{}, err
		}

		if next == '/' {
			return Token{}, discardSingleLineComment(r)
		}

		if next == '*' {
			return Token{}, discardMultilineComment(r)
		}

		if err = r.UnreadRune(); err != nil {
			return Token{}, err
		}
	}

	return Token{}, fmt.Errorf("syntax error: unexpected token '%c'", init)
}

func discardWhitespace(r io.RuneScanner) (init rune, err error) {
	for {
		if init, _, err = r.ReadRune(); err != nil {
			return
		}

		if !unicode.IsSpace(init) {
			return
		}
	}
}

func readInt(r io.RuneScanner, init rune) (Token, error) {
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

func readString(r io.RuneScanner, quote rune) (Token, error) {
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

func discardSingleLineComment(r io.RuneReader) error {
	for {
		char, _, err := r.ReadRune()
		if err != nil {
			return err
		}

		if char == '\n' {
			return ErrSkip
		}
	}
}

func discardMultilineComment(r io.RuneReader) error {
	var star bool
	for {
		char, _, err := r.ReadRune()
		if err != nil {
			return err
		}

		if star = char == '*'; star && char == '/' {
			return ErrSkip
		}
	}
}

type Token struct {
	Kind  Kind
	Value string
}
