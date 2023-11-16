package utils

import (
	"testing"
)

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}
func TestUUID(t *testing.T) {
	t.Log(UUID())
}
func BenchmarkSUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SUID()
	}
}
func TestSUID(t *testing.T) {
	t.Log(SUID())
}

func BenchmarkRUID16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RUID(16)
	}
}

func TestRUID16(t *testing.T) {
	t.Log(RUID(16))
}

func BenchmarkRUID8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RUID(8)
	}
}

func TestRUID8(t *testing.T) {
	t.Log(RUID(8))
}
func BenchmarkDUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DUID(8)
	}
}
func TestDUID(t *testing.T) {
	t.Log(DUID(8))
}
