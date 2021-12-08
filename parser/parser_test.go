package parser_test

import (
	"strings"
	"testing"

	"github.com/lthibault/eva/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	n, err := parser.New(strings.NewReader("42")).Parse()
	assert.NoError(t, err)
	assert.Equal(t, parser.Int, n.Kind())
	assert.Equal(t, parser.IntLiteral{42}, n)
}

func TestParseLiteral(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		src  string
		kind parser.Kind
		want parser.Literal
	}{{
		src:  "42",
		kind: parser.Int,
		want: parser.IntLiteral{42},
	}, {
		src:  `"test"`,
		kind: parser.String,
		want: parser.StringLiteral{"test"},
	}, {
		src:  "'test'",
		kind: parser.String,
		want: parser.StringLiteral{"test"},
	}} {
		n, err := parser.New(strings.NewReader(tt.src)).ParseLiteral()
		assert.NoError(t, err)
		assert.Equal(t, tt.kind, n.Kind())
		assert.Equal(t, tt.want, n)
	}
}

func TestParseInt(t *testing.T) {
	t.Parallel()

	n, err := parser.New(strings.NewReader("42")).ParseInt()
	assert.NoError(t, err)
	assert.Equal(t, parser.Int, n.Kind())
	assert.Equal(t, 42, n.Value, "value should be '42'")
}

func TestParseString(t *testing.T) {
	t.Parallel()
	t.Helper()

	t.Run("DoubleQuote", func(t *testing.T) {
		t.Parallel()

		n, err := parser.New(strings.NewReader(`"Hello, Eva!"`)).ParseString()
		assert.NoError(t, err)
		assert.Equal(t, parser.String, n.Kind())
		assert.Equal(t, "Hello, Eva!", n.Value, "value should be 'Hello, Eva'")
	})

	t.Run("SingleQuote", func(t *testing.T) {
		t.Parallel()

		n, err := parser.New(strings.NewReader("'Hello, Eva!'")).ParseString()
		assert.NoError(t, err)
		assert.Equal(t, parser.String, n.Kind())
		assert.Equal(t, "Hello, Eva!", n.Value, "value should be 'Hello, Eva'")
	})
}
