package hash

import (
	"testing"
)

func BenchmarkAppendHash(b *testing.B) {
	hash := New(1 << 14)

	for i := 0; i < b.N; i++ {
		hash.Append(16, 15)
	}
}

func BenchmarkGetValue(b *testing.B) {
	hash := New(1 << 14)

	for i := 0; i < b.N; i++ {
		hash.GetValue(16)
	}
}

func Test(t *testing.T) {
	hash := New(1 << 14)

	hash.Append(16, 100)
	value, _ := hash.GetValue(16)

	if value != 100 {
		t.Error("wrong")
	}
}
