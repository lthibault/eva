package eva

import "unsafe"

const (
	Int32 = iota
	String
)

type Type uint8

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
	v = *(*Value)(unsafe.Pointer(&s))
	return v
}

func (v Value) IsPointer() bool { return uintptr(v.ptr) != 0 }

func (v Value) Int32() int32   { return int32(v.val) }
func (v Value) String() string { return *(*string)(unsafe.Pointer(&v)) }
