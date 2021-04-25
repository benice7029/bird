package bird

import "testing"

func TestBird(t *testing.T) {
    want := "Bird chu chu!"
    if got := Bark(); got != want {
        t.Errorf("Bark() = %q, want %q", got, want)
    }
}