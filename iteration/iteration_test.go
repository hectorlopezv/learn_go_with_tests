package iteration

import (
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func correctAssesment(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected '%s' but got '%s'", got, want)
	}
}
func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		correctAssesment(t, repeated, expected)

	})

	t.Run("repeat N times, if N is passed as a argument", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"
		correctAssesment(t, repeated, expected)
	})
}
