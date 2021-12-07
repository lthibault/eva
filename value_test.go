package eva_test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/lthibault/eva"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValue_IsPointer(t *testing.T) {
	t.Parallel()

	assert.False(t, eva.Value{}.IsPointer(),
		"eva.Value should default to value type")

	assert.False(t, eva.NewInt32(42).IsPointer(),
		"int32 should be a value type")

	assert.False(t, eva.NewString("").IsPointer(),
		"string should be a pointer type")
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
	t.Helper()

	t.Run("Alloc", func(t *testing.T) {
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
	})

	t.Run("Add", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			message    string
			arg0, arg1 eva.Value
			want       int32
			err        string
		}{{
			message: "mismatched types should fail",
			arg0:    eva.NewInt32(42),
			arg1:    eva.NewString("fail"),
		}, {
			message: "should successfully add two int32 types",
			arg0:    eva.NewInt32(42),
			arg1:    eva.NewInt32(3),
			want:    45,
		}} {
			got, err := eva.Add(tt.arg0, tt.arg1)
			if tt.err != "" {
				assert.EqualError(t, err, tt.err)
			} else {
				assert.Equal(t, tt.want, got.Int32())
			}
		}
	})

	t.Run("Mul", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			message    string
			arg0, arg1 eva.Value
			want       int32
			wantErr    bool
		}{{
			message: "mismatched types should fail",
			arg0:    eva.NewInt32(42),
			arg1:    eva.NewString("fail"),
			wantErr: true,
		}, {
			message: "should successfully multiply two int32 types",
			arg0:    eva.NewInt32(21),
			arg1:    eva.NewInt32(2),
			want:    42,
		}} {
			got, err := eva.Mul(tt.arg0, tt.arg1)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.want, got.Int32())
				}
			}
		}
	})
}

func BenchmarkValue_Int32(b *testing.B) {
	b.Helper()

	b.Run("Alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if eva.NewInt32(int32(i)).Int32() != int32(i) {
				panic(fmt.Sprintf("expected %d, got %d", i, eva.NewInt32(int32(i)).Int32()))
			}
		}
	})

	b.Run("Add", func(b *testing.B) {
		var (
			arg0 = eva.NewInt32(42)
			arg1 = eva.NewInt32(3)
		)

		b.ResetTimer()
		if _, err := eva.Add(arg0, arg1); err != nil {
			panic(err)
		}
	})

	b.Run("Mul", func(b *testing.B) {
		var (
			arg0 = eva.NewInt32(42)
			arg1 = eva.NewInt32(3)
		)

		b.ResetTimer()
		if _, err := eva.Mul(arg0, arg1); err != nil {
			panic(err)
		}
	})
}

func TestValue_String(t *testing.T) {
	t.Parallel()
	t.Helper()

	t.Run("Alloc", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			message string
			s       string
		}{{
			message: "empty string should be valid",
		}, {
			message: "non-empty string should be valid",
			s:       "hello, eva!",
		}} {
			assert.Equal(t, tt.s, eva.NewString(tt.s).String(), tt.message)
		}
	})

	t.Run("Add", func(t *testing.T) {
		t.Parallel()

		v, err := eva.Add(eva.NewString("foo"), eva.NewString("bar"))
		require.NoError(t, err)
		assert.Equal(t, "foobar", v.String())
	})

	t.Run("Mul", func(t *testing.T) {
		t.Parallel()

		var b strings.Builder
		for i := 0; i < 100; i++ {
			b.WriteString("foobar")
		}

		for _, tt := range []struct {
			message    string
			arg0, arg1 eva.Value
			want       string
			wantErr    bool
		}{{
			message: "should fail if first argument is not of type 'string'",
			arg0:    eva.NewInt32(42),
			arg1:    eva.NewString("fail"),
			wantErr: true,
		}, {
			message: "should successfully multiply string by int32",
			arg0:    eva.NewString("foobar"),
			arg1:    eva.NewInt32(100),
			want:    b.String(),
		}} {
			got, err := eva.Mul(tt.arg0, tt.arg1)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got.String())
			}
		}
	})
}

func BenchmarkValue_String(b *testing.B) {
	b.Helper()

	b.Run("Alloc", func(b *testing.B) {
		var (
			s  = "hello, eva!"
			es = eva.NewString(s)
		)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = es.String()
		}
	})

	b.Run("Add", func(b *testing.B) {
		var (
			arg0 = eva.NewString("foo")
			arg1 = eva.NewString("bar")
		)

		b.ResetTimer()
		if _, err := eva.Add(arg0, arg1); err != nil {
			panic(err)
		}
	})

	b.Run("Mul", func(b *testing.B) {
		var (
			arg0 = eva.NewString("foo")
			arg1 = eva.NewInt32(100)
		)

		b.ResetTimer()
		if _, err := eva.Mul(arg0, arg1); err != nil {
			panic(err)
		}
	})
}
