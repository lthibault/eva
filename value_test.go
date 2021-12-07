package eva_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/lthibault/eva"
	"github.com/stretchr/testify/assert"
)

func TestValue_IsPointer(t *testing.T) {
	t.Parallel()

	assert.False(t, eva.Value{}.IsPointer(),
		"eva.Value should default to value type")

	assert.False(t, eva.NewInt32(42).IsPointer(),
		"int32 should be a value type")
}

func BenchmarkValue_IsPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// We just use int32 for benchmarks.  There should be no difference
		// in performance in cases where uintptr(v.ptr) != 0.
		if eva.NewInt32(int32(i)).IsPointer() == true {
			panic("int32 should not be pointer type")
		}
	}
}

func TestValue_Int32(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		message string
		i       int32
	}{{
		message: "zero should be valid",
	}, {
		message: "max int32 should be valid",
		i:       math.MaxInt32,
	}, {
		message: "min int32 should be valid",
		i:       math.MinInt32,
	}} {
		assert.Equal(t, tt.i, eva.NewInt32(tt.i).Int32(), tt.message)
	}
}

func BenchmarkValue_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if eva.NewInt32(int32(i)).Int32() != int32(i) {
			panic(fmt.Sprintf("expected %d, got %d", i, eva.NewInt32(int32(i)).Int32()))
		}
	}
}
