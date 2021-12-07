package eva

import (
	"fmt"
	"strings"
	"unsafe"
)

const (
	Int32 = iota
	String
)

type Type uint8

func (t Type) String() string {
	switch t {
	case Int32:
		return "int32"
	case String:
		return "string"
	default:
		panic(fmt.Sprintf("invalid type '%0x'", uint8(t)))
	}
}

type Value struct {
	// The first two fields have a memory representation equivalent
	// to the following native Go types:
	//  - string
	//  - []byte
	ptr unsafe.Pointer
	val uint64

	// TODO(performance):  align?
	Type Type
}

func NewInt32(i int32) Value {
	return Value{
		Type: Int32,
		val:  uint64(i),
	}
}

func NewString(s string) Value {
	v := Value{Type: String}
	*(*string)(unsafe.Pointer(&v)) = s
	return v
}

func (v Value) IsPointer() bool { return uintptr(v.ptr) != 0 }

func (v Value) Int32() int32   { return int32(v.val) }
func (v Value) String() string { return *(*string)(unsafe.Pointer(&v)) }

func Add(v0, v1 Value) (Value, error) {
	// same type ?
	if v0.Type^v1.Type == 0 {
		switch v0.Type {
		case String:
			return NewString(v0.String() + v1.String()), nil

		case Int32:
			return NewInt32(v0.Int32() + v1.Int32()), nil
		}
	}

	return Value{}, fmt.Errorf("cannot add type '%s' to '%s'", v0.Type, v1.Type)
}

func Mul(v0, v1 Value) (Value, error) {
	if v1.Type == Int32 {
		switch v0.Type {
		case Int32:
			return NewInt32(v0.Int32() * v1.Int32()), nil

		case String:
			var (
				b strings.Builder
				s = v0.String()
				i = v1.Int32()
			)

			if i > 0 {
				for b.Grow(int(i)); i > 0; i-- {
					b.WriteString(s)
				}
			}

			return NewString(b.String()), nil
		}
	}

	return Value{}, fmt.Errorf("cannot multiply type '%s' with '%s'", v0.Type, v1.Type)
}
