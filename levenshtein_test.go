package levenshtein

import (
	"testing"
)

const (
	sA = "わがはいは犬である。"
	sB = "わがはいは猫である。"
	sC = "わがはいは猫ではない。"
)

func TestDistance(t *testing.T) {
	r1 := []rune(sB)
	r2 := []rune(sC)
	if d := Distance(r1, r2); d != 3 {
		t.Errorf("distance not equal: %d", d)
	}
}

func TestPercent(t *testing.T) {
	r1 := []rune(sA)
	r2 := []rune(sB)
	if p := Percent(r1, r2); p != 90.0 {
		t.Errorf("percent not equal: %f", p)
	}
}

func BenchmarkDistance(b *testing.B) {
	r1 := []rune(sA)
	r2 := []rune(sB)
	for i := 0; i < b.N; i++ {
		_ = Distance(r1, r2)
	}
}

func BenchmarkPercent(b *testing.B) {
	r1 := []rune(sA)
	r2 := []rune(sB)
	for i := 0; i < b.N; i++ {
		_ = Percent(r1, r2)
	}
}
