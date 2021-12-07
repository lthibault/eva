package eva

import "unsafe"

const (
	Int32 = iota
)

type Type uint8

type Value struct {
	Type Type
	val  uint64
	ptr  unsafe.Pointer
}

func NewInt32(i int32) Value {
	return Value{
		Type: Int32,
		val:  uint64(i),
	}
}

func (v Value) IsPointer() bool { return uintptr(v.ptr) != 0 }

func (v Value) Int32() int32 { return int32(v.val) }
