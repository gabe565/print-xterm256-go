package main

import "testing"

func TestGenerate(t *testing.T) {
	v := Generate()
	if len(v) > bufCap {
		t.Errorf("Generate() got len = %d, want <= %d", len(v), bufCap)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for b.Loop() {
		_ = Generate()
	}
}
