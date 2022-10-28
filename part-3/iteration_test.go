package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 0)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

// windows powershell
// go test -bench="."
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
