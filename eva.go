// Package eva implements the Eva virtual machine.
package eva

import (
	"bytes"
	"fmt"
	"io"
)

const stackLimit = 4096

type VM struct {
	constants vector
	stack     [stackLimit]Value
	sp        int // stack pointer
}

// Exec the program.
func (vm *VM) Exec(program string) (Value, error) {
	// 1. parse the program

	// 2. compile to bytecode

	vm.constants.PushBack(NewString("Hello, Eva!"))

	var code = []byte{
		OP_CONST, 0, // const pool index
		OP_HALT,
	}

	// 3. evaluate the bytecode
	return vm.Eval(bytes.NewBuffer(code))
}

func (vm *VM) Eval(r io.ByteReader) (Value, error) {
	for {
		op, err := r.ReadByte()
		if err != nil {
			return Value{}, err
		}

		switch op {
		case OP_HALT:
			return vm.pop(), nil

		case OP_CONST:
			c, err := vm.getConst(r)
			if err != nil {
				return Value{}, err
			}
			vm.push(c) // push the result of the 'const' operation to the stack

		case OP_ADD, OP_SUB, OP_MUL, OP_DIV:
			arg1 := vm.pop() // n.b.:  stack is FILO
			arg0 := vm.pop()
			vm.push(binaryOp(op).Eval(arg0, arg1))

		default:
			// NOTE:  when refactoring to array/index-based opcodes, a non-existant
			//        op will likewise cause a panic (array out-of-bounds).
			panic(fmt.Sprintf("unknown opcode %x", op))
		}
	}
}

// push a value onto the stack
func (vm *VM) push(v Value) {
	if vm.sp == stackLimit-1 {
		panic("stack overflow")
	}
	vm.stack[vm.sp] = v
	vm.sp++
}

// pop a value from the stack
func (vm *VM) pop() Value {
	if vm.sp == 0 {
		panic("pop from empty stack")
	}
	// sp points to next free slot; decr then return value at index.
	vm.sp--
	return vm.stack[vm.sp]
}

func (vm *VM) getConst(r io.ByteReader) (Value, error) {
	cidx, err := r.ReadByte()
	return vm.constants[cidx], err
}

type binaryOp uint8

func (op binaryOp) Eval(arg0, arg1 Value) Value {
	// TODO(performance):  use array-indexing trick to speed this up.
	//                     We chan right-shift (>>) the opcode by some fixed
	//                     amount to turn it into an array index.

	// TODO:  make this work with other types.
	switch uint8(op) {
	case OP_ADD:
		return NewInt32(arg0.Int32() + arg1.Int32())

	case OP_SUB:
		return NewInt32(arg0.Int32() - arg1.Int32())

	case OP_MUL:
		return NewInt32(arg0.Int32() * arg1.Int32())

	case OP_DIV:
		return NewInt32(arg0.Int32() / arg1.Int32())

	default:
		panic(fmt.Sprintf("unrecognized binary operation '%0x'", op))
	}
}

type vector []Value

func (vs *vector) PushBack(v Value) { *vs = append(*vs, v) }
