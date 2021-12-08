package parser_test

import (
	"io"
	"strings"
	"testing"

	"github.com/lthibault/eva/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	n, err := parser.Parse(strings.NewReader(`42`))
	assert.ErrorIs(t, err, io.EOF, "error should be EOF")
	assert.Equal(t, "42", n.Value, "value should be '\"42\"'")
	assert.Equal(t, "NumericLiteral", n.Type())
}

func TestParseNumericLiteral(t *testing.T) {
	t.Parallel()

	n, err := parser.Parse(strings.NewReader(`42`))
	assert.ErrorIs(t, err, io.EOF, "error should be EOF")
	assert.Equal(t, "42", n.Value, "value should be '\"42\"'")
	assert.Equal(t, "NumericLiteral", n.Type())
}
