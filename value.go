package eva

import "unsafe"

const (
	NUMBER = iota
)

type Type uint8

type Value struct {
	Type  Type
	Value unsafe.Pointer // TODO:  tagged pointers?
}

func NewNumber(i int) *Value {
	return &Value{
		Type:  NUMBER,
		Value: unsafe.Pointer(&i),
	}
}

func (v Value) Number() int { return *(*int)(v.Value) }
