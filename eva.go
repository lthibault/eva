// Package eva implements the Eva virtual machine.
package eva

import (
	"bytes"
	"io"
)

type VM struct{}

// Exec the program.
func (vm VM) Exec(program string) error {
	// 1. parse the program

	// 2. compile to bytecode

	var code = []byte{
		OP_HALT,
	}

	// 3. evaluate the bytecode
	return vm.Eval(bytes.NewBuffer(code))
}

func (vm VM) Eval(r io.ByteReader) error {
	for {
		op, err := r.ReadByte()
		if err != nil {
			return err
		}

		switch op {
		case OP_HALT:
			return nil
		}
	}
}
