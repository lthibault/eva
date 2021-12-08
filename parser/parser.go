package parser

import (
	"fmt"
	"io"
	"strconv"
)

func Parse(r io.RuneReader) (Literal, error) {
	return New(r).Parse()
}

type Parser struct {
	r    io.RuneReader
	next Token
	err  error
}

func New(r io.RuneReader) *Parser {
	p := &Parser{r: r}
	p.next, p.err = ReadToken(r)
	return p
}

// Parse the program.
//
// Program
//  : Literal
//  ;
//
func (p *Parser) Parse() (Literal, error) {
	return p.ParseLiteral()
}

// ParseLiteral consumes a literal.
//
// Literal
//  : NumericLiteral
//  : StringLiteral
//  ;
//
func (p *Parser) ParseLiteral() (Literal, error) {
	switch p.next.Kind {
	case Int:
		return p.ParseInt()

	case String:
		return p.ParseString()

	default:
		return nil, fmt.Errorf("syntax error: unexpected literal production")
	}
}

// ParseInt consumes an integer literal.
//
// NumericLiteral
//  : Int
//  : Float
//  : Complex
//  ;
//
func (p *Parser) ParseInt() (IntLiteral, error) {
	t, err := p.consume(Int)
	if err != nil {
		return IntLiteral{}, err
	}

	i, err := strconv.Atoi(t.Value)
	if err != nil {
		return IntLiteral{}, err
	}

	return IntLiteral{
		Value: i,
	}, err
}

// ParseString consumes a string literal.
//
// StringLiteral
//  : String
//  ;
//
func (p *Parser) ParseString() (s StringLiteral, err error) {
	var t Token
	if t, err = p.consume(String); err != nil {
		return
	}

	if t.Value != "" {
		s.Value = t.Value[1 : len(t.Value)-1]
	}

	return
}

func (p *Parser) consume(k Kind) (t Token, err error) {
	if err = p.err; err != nil {
		return
	}

	if t = p.next; t.Kind != k {
		err = fmt.Errorf("unexpected token type '%s', expected '%s'", t.Kind, k)
		return
	}

	p.next, p.err = ReadToken(p.r)
	return
}

type Literal interface {
	Kind() Kind
}

type IntLiteral struct {
	Value int
}

func (IntLiteral) Kind() Kind { return Int }

type StringLiteral struct {
	Value string
}

func (StringLiteral) Kind() Kind { return String }
