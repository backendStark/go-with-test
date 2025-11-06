package iterations

import "testing"

// func TestRepeat(t *testing.T) {
// 	got := Repeat("a", 5)
// 	expected := "aaaaa"

// 	if got != expected {
// 		t.Errorf("expected %q but got %q", expected, got)
// 	}
// }

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
