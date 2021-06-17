package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestIsSameString(t *testing.T) {
	got := IsSameString("a", "a")
	want := true

	if got != want {
		t.Errorf("expected %t, but got %t", got, want)
	}
}
