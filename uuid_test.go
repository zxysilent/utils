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

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}
func TestUUID(t *testing.T) {
	t.Log(UUID())
}
func TestBits(t *testing.T) {
	t.Logf("%b\n", 253402246800000/16)
}
