package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("repeat character 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})
	t.Run("repeat character 8 times", func(t *testing.T) {
		got := Repeat("b", 8)
		want := "bbbbbbbb"

		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}