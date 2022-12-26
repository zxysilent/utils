package utils

import (
	"testing"
)

func BenchmarkSUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SUID()
	}
}
func TestSUID(t *testing.T) {
	t.Log(SUID())
}
