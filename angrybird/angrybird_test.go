package angrybird

import "testing"

func TestBird(t *testing.T) {
    want := "Angrybird aaaaaaaaaa!"
    if got := Bark(); got != want {
        t.Errorf("Bark() = %q, want %q", got, want)
    }
}