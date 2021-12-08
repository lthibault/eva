package parser

import "errors"

var (
	// ErrSkip causes the tokenizer to skip the current token.
	ErrSkip = errors.New("skip")
)
